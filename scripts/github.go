package scripts

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"

	"gashubsync/utils"
)

func CommitToGithub(repoURL string, content string) {
	ctx := context.Background()

	token := utils.ReadToken("./github_token.txt")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Parse the repo URL to extract the owner and repo
	splitURL := strings.Split(repoURL, "/")
	owner, repo := splitURL[len(splitURL)-2], splitURL[len(splitURL)-1]

	// Create a new file in the repo
	fileContent := []byte(content)
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("Initial commit from Google Apps Script"),
		Content:   fileContent,
		Committer: &github.CommitAuthor{Name: github.String("MyApp"), Email: github.String("myapp@example.com")},
	}

	_, _, err := client.Repositories.CreateFile(ctx, owner, repo, "script.gs", opts)
	if err != nil {
		log.Fatalf("Repositories.CreateFile returned error: %v", err)
	}
}
