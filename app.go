package main

import (
	"net/http"
)

func init() {
	initializeConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/fetch/regular", FetchRegular)
	mux.HandleFunc("/fetch/playoff", FetchPlayoff)
	mux.HandleFunc("/api/player/list", PlayerList)
	mux.HandleFunc("/api/playoffplayer/list", PlayoffPlayerList)
	http.Handle("/", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}
