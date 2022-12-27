package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // blank import of pprof
	// will make profiler available at localhost:3000/debug/pprof
)

func PprofServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
