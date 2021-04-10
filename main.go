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

	ops := []app.DeltaOps{
		{
			Kind: "MOV",
			Mov: app.OpsMov{
				FromIndex: 3,
				Length:    1,
				ToIndex:   1,
			},
		},
		{
			Kind: "MOV",
			Mov: app.OpsMov{
				FromIndex: 3,
				Length:    1,
				ToIndex:   2,
			},
		},
	}

	changeRes, err := app.PostRootListChanges(ops, res.Revision, token.AccessToken, userId)
	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("res: %+v\n", changeRes)
}
