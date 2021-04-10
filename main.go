package main

import (
	"fmt"
	webplayer "github.com/mirrorfm/spotify-webplayer-token/app"
	"github.com/mirrorfm/unofficial-spotify-api/app"
	"os"
)

func main() {
	token, err := webplayer.GetAccessTokenFromEnv()
	if err != nil && !token.IsAnonymous {
		os.Exit(1)
	}

	userId, exists := app.GetUserIdFromEnv()
	if !exists {
		os.Exit(1)
	}

	res, err := app.GetRootList(token.AccessToken, userId)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("res: %+v\n", res)

	changeRes, err := app.PostRootListChanges(1, 3, res.Revision, token.AccessToken, userId)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("res: %+v\n", changeRes)
}
