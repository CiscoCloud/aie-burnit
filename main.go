package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/danielkrainas/aie-burnit/marathon"
	"github.com/danielkrainas/aie-burnit/names"
	"github.com/danielkrainas/aie-burnit/resources"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	LOCAL_APP_ID = "local-app"
)

var (
	MARATHON_APP_ID = ""
	instanceName    = ""
	alerts          = 0
	marathonClient  marathon.Client
)

type updateRequest struct {
	Resource string `json:"resource,omitempty"`
	Value    string `json:"value,omitempty"`
	Action   string `json:"action,omitempty"`
	Host     string `json:"host,omitempty"`
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	op := &updateRequest{}
	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(op); err != nil {
		fmt.Printf("error decoding update: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := strconv.ParseFloat(op.Value, 32)
	if err != nil {
		fmt.Printf("error parsing value: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if op.Host == "" || r.URL.Path == "/update/self" {
		switch op.Resource {
		case "memory":
			if op.Action == "reset" {
				resources.ResetMemoryUsage()
			} else {
				resources.SetMemoryUsage(value)
			}
		}

		w.WriteHeader(http.StatusNoContent)
	} else {
		content, err := json.Marshal(op)
		if err != nil {
			fmt.Printf("error reencoding update: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := http.Post(fmt.Sprintf("http://%s/update/self", op.Host), "application/json", strings.NewReader(string(content)))
		if err != nil {
			fmt.Printf("error relaying update: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp.Body.Close()
		w.WriteHeader(resp.StatusCode)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		alerts++
		w.WriteHeader(http.StatusOK)
	} else {
		assetHandler(w, r)
	}
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./assets/style.css")
		return
	} else if r.URL.Path == "/app.js" {
		http.ServeFile(w, r, "./assets/app.js")
		return
	} else if r.URL.Path == "/" {
		http.ServeFile(w, r, "./assets/index.html")
		return
	} else if r.URL.Path == "/cisco-logo-white.png" {
		http.ServeFile(w, r, "./assets/cisco-logo-white.png")
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf(`{
		"name": %q,
		"host": %q,
		"memory_usage": "%.2f"
	}`, instanceName, r.Host, resources.GetMemoryUsage()))
}

func aggregateStatusHandler(w http.ResponseWriter, r *http.Request) {
	app, err := marathonClient.GetApp(MARATHON_APP_ID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, fmt.Sprintf(`{ "errors": [%q] }`, err.Error()))
		return
	}

	if app == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	results := make([]string, 0)
	for _, t := range app.Tasks {
		status, err := getStatus(t.Host, t.Ports[0])
		if err == nil {
			results = append(results, string(status))
		}
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(results, ",")))
}

func getStatus(host string, port int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/status", host, port))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(s), nil
	}

	return "", err
}

func determineAppId() {
	instanceName = names.Generate()
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName != "" {
		fmt.Printf("svc name=%q\n", serviceName)
		hostVarName := strings.ToUpper(serviceName)
		hostVarName = "HOST_" + strings.Replace(hostVarName, "-", "_", -1)
		if os.Getenv(hostVarName) == "" {
			MARATHON_APP_ID == serviceName
		} else {
			u, err := url.Parse(os.Getenv(hostVarName))
			if err == nil {
				MARATHON_APP_ID = strings.Split(u.Host, ".")[0]
			}
		}
	} else {
		fmt.Println("SERVICE_NAME not found")
	}

	if MARATHON_APP_ID == "" {
		MARATHON_APP_ID = LOCAL_APP_ID
	}

	fmt.Printf("app=%s\n", MARATHON_APP_ID)
}

func setupMarathon() {
	var err error
	if os.Getenv("MOCK") != "" {
		fmt.Println("mocks enabled")
		marathonClient = marathon.NewMockClient()
		err = nil
	} else {
		marathonClient, err = marathon.NewClient()
	}

	if err != nil {
		panic(err)
	}
}

func main() {
	determineAppId()
	setupMarathon()
	app, err := marathonClient.GetApp(MARATHON_APP_ID)
	if err != nil {
		panic(err)
	} else if app == nil {
		panic("couldn't get app from marathon")
	}

	resources.SetMemoryLimit(float64(app.Mem))
	http.HandleFunc("/", http.HandlerFunc(defaultHandler))
	http.HandleFunc("/update", http.HandlerFunc(updateHandler))
	http.HandleFunc("/update/self", http.HandlerFunc(updateHandler))
	http.HandleFunc("/style.css", http.HandlerFunc(assetHandler))
	http.HandleFunc("/app.js", http.HandlerFunc(assetHandler))
	http.HandleFunc("/status", http.HandlerFunc(statusHandler))
	http.HandleFunc("/status/all", http.HandlerFunc(aggregateStatusHandler))
	fmt.Println("Example app listening at http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
