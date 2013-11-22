package main

import (
	"fmt"
	"github.com/tmiller/go-pivotal-tracker-api"
	"io/ioutil"
	"os"
	"strings"
)

var pivotalTracker pt.PivotalTracker

func main() {
	initPivotalTracker()
	printStory()
}

func printStory() {

	storyId := strings.TrimSpace(os.Args[1])

	if story, err := pivotalTracker.FindStory(storyId); err == nil {
		fmt.Printf("- %v\n  %v\n\n", story.Name, story.Url)
	} else {
		fmt.Println(err)
		os.Exit(2)
	}
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
