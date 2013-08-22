package main

import (
	"fmt"
	"github.com/tmiller/go-pivotal-tracker-api"
	"os"
	"strings"
)

var pivotalTracker pt.Pivotaltracker

func main() {
	initPivotalTracker()
}

func initPivotalTracker() {
	configFilePath := os.ExpandEnv("${HOME}/.pivotal_tracker_api_key")
	configFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		fmt.Printf("Please put your Pivotal Tracker API key in %s\n", configFilePath)
		os.Exit(1)
	}

	pivotalTrackerApiKey := strings.TrimSpace(string(configFile))
	pivotalTracker = pt.PivotalTracker{pivotalTrackerApiKey}
}
