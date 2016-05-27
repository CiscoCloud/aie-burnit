package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var nm sync.Mutex
var N int = 0
var delta int = 524288 // 512kb
var alerts = 0

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	nm.Lock()
	defer nm.Unlock()

	if r.Method == "POST" {
		alerts++
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Fprintf(w, `<head><meta http-equiv="refresh" content="2"></head><body>`)
		fmt.Fprintf(w, "<h1>hello from burnit: %.2f MB and %d notifications</h1>\n</body>", float32(N)/1048576.0, alerts)
	}
}

func main() {
	N = 0
	i := 0
	blob := make([][]byte, 100000)
	go func() {
		for {
			nm.Lock()
			blob[i] = bytes.Repeat([]byte{70}, delta)
			i++
			N += delta
			nm.Unlock()
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	http.HandleFunc("/", defaultHandler)
	fmt.Println("Example app listening at http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
