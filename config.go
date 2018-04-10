package main

import "os"

var MSFUsername = ""
var MSFPassword = ""

var ListPlayersMemcacheKey = "LIST-PLAYERS"

func initializeConfig() {
	MSFUsername = os.Getenv("MSF_USERNAME")
	MSFPassword = os.Getenv("MSF_PASSWORD")
}
