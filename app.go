package main

import (
	"fmt"
	"net/http"
)

func init() {
	initializeConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/fetch", Fetch)
	mux.HandleFunc("/api/player/list", PlayerList)
	http.Handle("/", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the app")
}
