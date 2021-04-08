package main

import (
	"fmt"
	webplayer "github.com/mirrorfm/spotify-webplayer-token/app"
	"github.com/mirrorfm/unofficial-spotify-api/app"
	"os"
)

func main() {
	token, err := webplayer.GetAccessTokenFromEnv()
	if err != nil {
		os.Exit(1)
	}

	cl := app.Spotify{
		Token: token,
	}

	err = cl.Reorder(0, 1)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(token)
}
