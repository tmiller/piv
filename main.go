package main

import (
	"flag"
	"fmt"
	"github.com/tmiller/go-pivotal-tracker-api"
	"io/ioutil"
	"os"
	"strings"
)

var pivotalTracker pt.PivotalTracker

func main() {
	flag.Parse()
	initPivotalTracker()
	printStory()
}

func printStory() {

	for _, storyId := range flag.Args() {

		if story, err := pivotalTracker.FindStory(storyId); err == nil {
			fmt.Printf("- %v\n  %v\n\n", story.Name, story.Url)
		} else {
			fmt.Println(err)
			os.Exit(2)
		}

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
