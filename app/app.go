package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const UserAgent = "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36"

func GetUserIdFromEnv() (string, bool) {
	return os.LookupEnv("SPOTIFY_USER_ID")
}

func RootList(token, userId string) (string, error) {
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

func Reorder(from, to int) error {
	return nil
}