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

	for _, storyId := range flag.Args() {
		if story, err := pivotalTracker.FindStory(storyId); err == nil {
			printStory(story)
		} else {
			printError(err)
		}
	}
}

func initPivotalTracker() {
	configFilePath := os.ExpandEnv("${HOME}/.pivotal_tracker_api_key")
	configFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		printError(err)
		os.Exit(1)
	}

	pivotalTrackerApiKey := strings.TrimSpace(string(configFile))
	pivotalTracker = pt.PivotalTracker{pivotalTrackerApiKey}
}

func printStory(story pt.Story) {
	fmt.Printf("- %v\n  %v\n\n", story.Name, story.Url)
}

func printError(err error) {
	errorMessage := fmt.Sprintf("%s\n", err.Error())
	os.Stderr.WriteString(errorMessage)
}
