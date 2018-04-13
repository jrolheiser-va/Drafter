package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

func PlayerList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(appengine.NewContext(r), 60*time.Second)
	defer cancel()
	var players []Player

	if _, err := memcache.JSON.Get(ctx, ListPlayersMemcacheKey, &players); err == memcache.ErrCacheMiss {
		q := datastore.NewQuery("Player")
		_, err = q.GetAll(ctx, &players)
		if err != nil {
			log.Errorf(ctx, "Error reading from Datastore: %v", err)
			return
		}
		if err = memcache.JSON.Set(ctx, &memcache.Item{Key: ListPlayersMemcacheKey, Object: players}); err != nil {
			log.Errorf(ctx, "Error setting item: %v", err)
			return
		}
	} else if err != nil {
		log.Errorf(ctx, "Error from memcache: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func PlayoffPlayerList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(appengine.NewContext(r), 60*time.Second)
	defer cancel()
	var players []Player

	if _, err := memcache.JSON.Get(ctx, ListPlayoffPlayersMemcacheKey, &players); err == memcache.ErrCacheMiss {
		q := datastore.NewQuery("PlayoffPlayer")
		_, err = q.GetAll(ctx, &players)
		if err != nil {
			log.Errorf(ctx, "Error reading from Datastore: %v", err)
			return
		}
		if err = memcache.JSON.Set(ctx, &memcache.Item{Key: ListPlayoffPlayersMemcacheKey, Object: players}); err != nil {
			log.Errorf(ctx, "Error setting item: %v", err)
			return
		}
	} else if err != nil {
		log.Errorf(ctx, "Error from memcache: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}
