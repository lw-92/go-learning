package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		sprintf := fmt.Sprintf("{\"time\":\"%s\"}", now)
		w.Write([]byte(sprintf))

	})
	http.ListenAndServe(":8080", nil)
}
