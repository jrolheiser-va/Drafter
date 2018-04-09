package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func PlayerList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(appengine.NewContext(r), 60*time.Second)
	defer cancel()
	q := datastore.NewQuery("Player")
	var players []Player
	_, err := q.GetAll(ctx, &players)
	if err != nil {
		log.Errorf(ctx, "fetching next Person: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}
