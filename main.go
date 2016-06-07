package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var nm sync.Mutex
var N int = 0
var delta int = 524288 // 512kb
var alerts = 0
var mode = "stop"

func updateHandler(w http.ResponseWriter, r *http.Request) {
	nm.Lock()
	defer nm.Unlock()
	r.ParseForm()

	action := r.Form.Get("action")
	fmt.Printf("updating: action=%q\n", action)
	if action != "" {
		mode = action
	} else {
		memoryStr := r.Form.Get("increment")
		if memory, err := strconv.ParseFloat(memoryStr, 32); err == nil {
			delta = int(memory * 1048576.0)
			fmt.Printf("updating increment to %d\n", delta)
		} else {
			fmt.Errorf("error: %s: %v\n", memoryStr, err)
		}
	}

	http.Redirect(w, r, "/", 302)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	nm.Lock()
	defer nm.Unlock()

	if r.Method == "POST" {
		alerts++
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Fprint(w, `<head></head><body>`)
		fmt.Fprintf(w, "<h1>hello from burnit: using ~%.2f MB and %d notifications</h1>\n", float32(N)/1048576.0, alerts)
		fmt.Fprint(w, `<form method="post" action="update">`)
		fmt.Fprintf(w, `<label>Increment Usage (MB)</label><input type="text" name="increment" value="%.2f" onblur="resume()" onfocus="pause()" />`, float32(delta)/1048576.0)
		fmt.Fprint(w, `<button type="submit">Update</button><br />`)
		if mode != "run" && mode != "pause" {
			fmt.Fprint(w, `<button name="action" type="submit" value="run">Start</button>`)
		} else if mode == "pause" {
			fmt.Fprint(w, `<button name="action" type="submit" value="run">Resume</button>`)
		} else if mode == "run" {
			fmt.Fprint(w, `<button name="action" type="submit" value="pause">Pause</button>`)
			fmt.Fprint(w, `<button name="action" type="submit" value="stop">Stop</button>`)
		}

		fmt.Fprint(w, `</form>`)
		fmt.Fprint(w, `<script>var refreshTimer = null;resume();`)
		fmt.Fprint(w, `function pause(){ clearTimeout(refreshTimer);refreshTimer = null; }`)
		fmt.Fprint(w, `function resume(){ refreshTimer = setTimeout(function () { window.location.href = window.location.href; }, 3000) }</script>`)
		fmt.Fprint(w, `<script src="https://code.jquery.com/jquery-2.2.4.min.js" integrity="sha256-BbhdlvQf/xTY9gja0Dq3HiwQF8LaCRTXxZKRutelT44=" crossorigin="anonymous"></script></body>`)
	}
}

func main() {
	N = 0
	i := 0
	blob := make([][]byte, 100000)
	go func() {
		for {
			nm.Lock()
			if mode == "run" {
				blob[i] = bytes.Repeat([]byte{70}, delta)
				i++
				N += delta
			} else if mode == "stop" {
				mode = "pause"
				blob = make([][]byte, 100000)
				alerts = 0
				N = 0
			}

			nm.Unlock()
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/update", updateHandler)
	fmt.Println("Example app listening at http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
