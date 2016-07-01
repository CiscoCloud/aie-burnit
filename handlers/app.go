package handlers

import ()

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resourceType := r.Form.Get("resource")
	switch resourceType {
	case "memory":
		if r.Form.Get("action") == "reset" {
			resources.ResetMemoryUsage()
		} else {
			memoryStr := r.Form.Get("memory-usage")
			if memory, err := strconv.ParseFloat(memoryStr, 32); err == nil {
				i := int(memory * MEGABYTE)
				fmt.Printf("updating memory usage to %d\n", i)
				resources.SetMemoryUsage(i)
			}
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func proxyUpdateHandler(w http.ResponseWriter, r *http.Request) {

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
	}`, instanceName, r.Host, resources.GetMemoryUsage()/MEGABYTE))
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
	results := make([]string, len(app.Tasks))
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
	fmt.Printf("fetching status from %s:%s", host, port)
	resp, err := http.Get(fmt.Sprintf("%s:%d/status", host, port))
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
