package main

import (
	"net/http"
)

func init() {
	initializeConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/fetch", Fetch)
	mux.HandleFunc("/api/player/list", PlayerList)
	http.Handle("/", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}
