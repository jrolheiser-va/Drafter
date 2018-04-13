package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"google.golang.org/appengine/urlfetch"
)

const cumulative_regular_stats = "https://api.mysportsfeeds.com/v1.2/pull/nhl/2017-2018-regular/cumulative_player_stats.json"
const cumulative_playoff_stats = "https://api.mysportsfeeds.com/v1.2/pull/nhl/2018-playoff/cumulative_player_stats.json"

func FetchRegular(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(appengine.NewContext(r), 60*time.Second)
	defer cancel()
	client := urlfetch.Client(ctx)
	req, err := http.NewRequest("GET", cumulative_regular_stats, nil)
	req.SetBasicAuth(MSFUsername, MSFPassword)
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		target := MSFCumulativePlayerStatsResponse{}
		json.NewDecoder(res.Body).Decode(&target)

		log.Infof(ctx, fmt.Sprintf("%v", target.CumulativePlayerStats.LastUpdatedOn))
		for _, msfPlayer := range target.CumulativePlayerStats.Players {
			if msfPlayer.PlayerStats.Stats.Points.Value == 0 {
				continue
			}
			playerKey := datastore.NewKey(ctx, "Player", msfPlayer.Description.ID, 0, nil)
			player := toPlayer(msfPlayer)
			if _, err := datastore.Put(ctx, playerKey, &player); err != nil {
				log.Errorf(ctx, "datastore.Put: %v", err)
			}
		}
		if err = memcache.Delete(ctx, ListPlayersMemcacheKey); err != nil {
			log.Warningf(ctx, "Error unsetting key: %v", err)
		}
	} else {
		log.Errorf(ctx, fmt.Sprintf("Recieved Error %d", res.StatusCode))
	}
	fmt.Fprintf(w, "HTTP GET returned status %v", res.Status)
}

func FetchPlayoff(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(appengine.NewContext(r), 60*time.Second)
	defer cancel()
	client := urlfetch.Client(ctx)
	req, err := http.NewRequest("GET", cumulative_playoff_stats, nil)
	req.SetBasicAuth(MSFUsername, MSFPassword)
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		target := MSFCumulativePlayerStatsResponse{}
		json.NewDecoder(res.Body).Decode(&target)

		log.Infof(ctx, fmt.Sprintf("%v", target.CumulativePlayerStats.LastUpdatedOn))
		for _, msfPlayer := range target.CumulativePlayerStats.Players {
			if msfPlayer.PlayerStats.Stats.Points.Value == 0 {
				continue
			}
			playerKey := datastore.NewKey(ctx, "PlayoffPlayer", msfPlayer.Description.ID, 0, nil)
			player := toPlayer(msfPlayer)
			if _, err := datastore.Put(ctx, playerKey, &player); err != nil {
				log.Errorf(ctx, "datastore.Put: %v", err)
			}
		}
		if err = memcache.Delete(ctx, ListPlayoffPlayersMemcacheKey); err != nil {
			log.Warningf(ctx, "Error unsetting key: %v", err)
		}
	} else {
		log.Errorf(ctx, fmt.Sprintf("Recieved Error %d", res.StatusCode))
	}
	fmt.Fprintf(w, "HTTP GET returned status %v", res.Status)
}

type MSFCumulativePlayerStatsResponse struct {
	CumulativePlayerStats MSFCumulativePlayerStats `json:"cumulativeplayerstats"`
}

type MSFCumulativePlayerStats struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	Players       []MSFPlayer `json:"playerstatsentry"`
}

type MSFPlayer struct {
	Description MSFDescription `json:"player"`
	Team        MSFTeam        `json:"team"`
	PlayerStats MSFPlayerStats `json:"stats"`
}

type MSFDescription struct {
	ID           string `json:"ID"`
	LastName     string `json:"LastName"`
	FirstName    string `json:"FirstName"`
	JerseyNumber string `json:"JerseyNumber"`
	Position     string `json:"Position"`
	Height       string `json:"Height"`
	Weight       string `json:"Weight"`
	BirthDate    string `json:"BirthDate"`
	Age          int64  `json:"Age,string"`
	BirthCity    string `json:"BirthCity"`
	BirthCountry string `json:"BirthCountry"`
	IsRookie     bool   `json:"IsRookie"`
}

type MSFTeam struct {
	ID           string `json:"ID"`
	City         string `json:"City"`
	Name         string `json:"Name"`
	Abbreviation string `json:"Abbreviation"`
}

type MSFPlayerStats struct {
	GamesPlayed MSFStatsDimension `json:"GamesPlayed"`
	Stats       MSFStats          `json:"stats"`
}

type MSFStats struct {
	Goals              MSFStatsDimension `json:"Goals"`
	Assists            MSFStatsDimension `json:"Assists"`
	Points             MSFStatsDimension `json:"Points"`
	HatTricks          MSFStatsDimension `json:"HatTricks"`
	PlusMinus          MSFStatsDimension `json:"PlusMinus"`
	Shots              MSFStatsDimension `json:"Shots"`
	BlockedShots       MSFStatsDimension `json:"BlockedShots"`
	ShotPercentage     MSFStatsDimension `json:"ShotPercentage"`
	Penalties          MSFStatsDimension `json:"Penalties"`
	PenaltyMinutes     MSFStatsDimension `json:"PenaltyMinutes"`
	PowerplayGoals     MSFStatsDimension `json:"PowerplayGoals"`
	PowerplayAssists   MSFStatsDimension `json:"PowerplayAssists"`
	PowerplayPoints    MSFStatsDimension `json:"PowerplayPoints"`
	ShorthandedGoals   MSFStatsDimension `json:"ShorthandedGoals"`
	ShorthandedAssists MSFStatsDimension `json:"ShorthandedAssists"`
	ShorthandedPoints  MSFStatsDimension `json:"ShorthandedPoints"`
	GameWinningGoals   MSFStatsDimension `json:"GameWinningGoals"`
	GameTyingGoals     MSFStatsDimension `json:"GameTyingGoals"`
	Hits               MSFStatsDimension `json:"Hits"`
	Faceoffs           MSFStatsDimension `json:"Faceoffs"`
	FaceoffWins        MSFStatsDimension `json:"FaceoffWins"`
	FaceoffLosses      MSFStatsDimension `json:"FaceoffLosses"`
	FaceoffPercent     MSFStatsDimension `json:"FaceoffPercent"`
}

type MSFStatsDimension struct {
	Dimension string `json:"@abbreviation"`
	Value     int64  `json:"#text,string"`
}
