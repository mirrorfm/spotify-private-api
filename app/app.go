package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetUserIdFromEnv() (string, bool) {
	return os.LookupEnv("SPOTIFY_USER_ID")
}

func GetRootList(token, userId string) (*RootListResponse, int, error) {
	url := fmt.Sprintf("https://spclient.wg.spotify.com/playlist/v2/user/%s/rootlist?decorate=revision%%2Clength%%2Cattributes%%2Ctimestamp%%2Cowner", userId)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error during GET request")
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return nil, 0, err
	}
	_ = resp.Body.Close()

	data := RootListResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("could not unmarshal JSON body")
		return nil, 0, err
	}

	return &data, resp.StatusCode, nil
}

func PostRootListChanges(ops []DeltaOps, baseRevision, token, userId string) (*RootListChangeResponse, int, error) {
	str := &ChangesPayload{
		BaseRevision: baseRevision,
		Deltas: []ChangeDelta{
			{
				Ops: ops,
			},
		},
	}

	jsonStr, err := json.Marshal(str)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, 0, err
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
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read request response body")
		return nil, 0, err
	}
	_ = resp.Body.Close()

	data := RootListChangeResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("could not unmarshal JSON body")
		return nil, 0, err
	}

	return &data, resp.StatusCode, nil
}
