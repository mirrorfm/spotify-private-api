package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const UserAgent = "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"

type OpsMov struct {
	FromIndex int `json:"fromIndex"`
	Length    int `json:"length"`
	ToIndex   int `json:"toIndex"`
}

type DeltaOps struct {
	Kind string `json:"kind"`
	Mov  OpsMov `json:"mov"`
}

type ChangeDelta struct {
	Ops []DeltaOps `json:"ops"`
}

type ChangesPayload struct {
	BaseRevision string        `json:"baseRevision"`
	Deltas       []ChangeDelta `json:"deltas"`
}

func GetUserIdFromEnv() (string, bool) {
	return os.LookupEnv("SPOTIFY_USER_ID")
}

func GetRootList(token, userId string) (string, error) {
	url := fmt.Sprintf("https://spclient.wg.spotify.com/playlist/v2/user/%s/rootlist?decorate=revision%%2Clength%%2Cattributes%%2Ctimestamp%%2Cowner", userId)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("user-agent", UserAgent)
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during GET request")
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return "", err
	}
	_ = resp.Body.Close()

	res := string(body)

	return res, nil
}

func PostRootListChanges(fromIndex, toIndex int, token, userId string) (string, error) {
	str := &ChangesPayload{
		BaseRevision: "AAAEIonITCt/8fIQgHqgSSUB2IdxrHIL",
		Deltas: []ChangeDelta{
			{
				Ops: []DeltaOps{
					{
						Kind: "MOV",
						Mov: OpsMov{
							fromIndex,
							1,
							toIndex,
						},
					},
				},
			},
		},
	}

	jsonStr, err := json.Marshal(str)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	url := fmt.Sprintf("https://spclient.wg.spotify.com/playlist/v2/user/%s/rootlist/changes", userId)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during POST request")
		return "", err
	}
	fmt.Printf("res: %+v\n", resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return "", err
	}
	_ = resp.Body.Close()

	res := string(body)

	return res, nil
}
