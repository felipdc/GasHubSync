package main

import (
	"fmt"
	"os"
)

func main() {
	// Take Google App Script URL and GitHub repo as command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: ghs <google-app-script-url> <github-repo-url>")
		os.Exit(1)
	}

	googleAppScriptURL := os.Args[1]
	githubRepoURL := os.Args[2]

	fmt.Println("Google Apps Script URL: ", googleAppScriptURL)
	fmt.Println("GitHub Repo URL: ", githubRepoURL)

	// TODO: Add code to interact with Google Drive API and GitHub API
}
