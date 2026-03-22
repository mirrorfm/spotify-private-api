# unofficial-spotify-api

Go client for Spotify's internal API endpoints not available in the [public Web API](https://developer.spotify.com/documentation/web-api).

Used by [mirror.fm](https://mirror.fm) to reorder playlists in a user's library.

## Usage

Requires a web player access token obtained via [spotify-webplayer-token](https://github.com/mirrorfm/spotify-webplayer-token) and a `SPOTIFY_USER_ID` env var.

```go
import (
    webPlayer "github.com/mirrorfm/spotify-webplayer-token/app"
    api "github.com/mirrorfm/unofficial-spotify-api/app"
)

func main() {
    token, _ := webPlayer.GetAccessTokenFromEnv()
    userId, _ := api.GetUserIdFromEnv()

    // Get all playlists in their current display order
    rootList, status, _ := api.GetRootList(token.AccessToken, userId)

    // Reorder playlists
    ops := api.DeltaOps{
        Kind: "MOV",
        Mov: api.OpsMov{FromIndex: 5, Length: 1, ToIndex: 0},
    }
    api.PostRootListChanges([]api.DeltaOps{ops}, rootList.Revision, token.AccessToken, userId)
}
```

## API endpoints

### `GET /playlist/v2/user/{userId}/rootlist`

Returns all playlists in their current display order with revision info.

### `POST /playlist/v2/user/{userId}/rootlist/changes`

Reorder playlists by submitting move operations against a base revision.
