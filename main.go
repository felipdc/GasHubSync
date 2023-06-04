package main

import (
	"fmt"
	"os"
	"time"

	"gashubsync/scripts"
	"gashubsync/utils"
)

func main() {
	// Take Google App Script ID and GitHub repo as command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: ghs <google-app-script-id> <github-repo-url>")
		os.Exit(1)
	}

	googleAppScriptID := os.Args[1]
	githubRepoURL := os.Args[2]

	fmt.Println("Google Apps Script ID: ", googleAppScriptID)
	fmt.Println("GitHub Repo URL: ", githubRepoURL)

	// Get the polling interval from config
	config := utils.ReadConfig("./config.json")
	pollingInterval := time.Duration(config.PollingIntervalSeconds) * time.Second

	// Get initial content of the App Script
	currentContent := scripts.GetAppScriptContent(googleAppScriptID)

	// Initial commit to GitHub
	scripts.CommitToGithub(githubRepoURL, currentContent)

	// Start polling at the specified interval
	for range time.Tick(pollingInterval) {
		newContent := scripts.GetAppScriptContent(googleAppScriptID)

		// If the contents are not equal, update GitHub
		if newContent != currentContent {
			scripts.CommitToGithub(githubRepoURL, newContent)
			currentContent = newContent
		}
	}
}
