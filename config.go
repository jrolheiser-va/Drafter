package main

import "os"

var MSFUsername = ""
var MSFPassword = ""

var ListPlayersMemcacheKey = "LIST-PLAYERS"
var ListPlayoffPlayersMemcacheKey = "LIST-PLAYOFF-PLAYERS"

func initializeConfig() {
	MSFUsername = os.Getenv("MSF_USERNAME")
	MSFPassword = os.Getenv("MSF_PASSWORD")
}
