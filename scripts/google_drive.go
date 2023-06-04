package scripts

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	drive "google.golang.org/api/drive/v3"

	"gashubsync/utils"
)

func GetAppScriptContent(fileId string) string {
	ctx := context.Background()

	token := utils.ReadToken("./google_token.txt")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	srv, err := drive.New(tc)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	file, err := srv.Files.Get(fileId).Download()
	if err != nil {
		log.Fatalf("Unable to retrieve file: %v", err)
	}

	defer file.Body.Close()
	data, err := ioutil.ReadAll(file.Body)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	return string(data)
}
