package github

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/stshontikidis/image_check/util"
)

// Dispatch will execute a repository_dispatch event against github repo
func Dispatch(repo string, token string) error {
	acceptType := "application/vnd.github.v3+json"
	url := "https://api.github.com/repos/" + repo + "/dispatches"

	body := []byte(`{"event_type": "base_update"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	util.CheckErr(err)

	req.Header.Set("Accept", acceptType)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	util.CheckErr(err)

	defer resp.Body.Close()

	if resp.StatusCode == 422 {
		fmt.Println("422")
		body, err := ioutil.ReadAll(resp.Body)
		util.CheckErr(err)

		return errors.New(string(body))

	} else if resp.StatusCode == 204 {
		fmt.Println("Success!")
	} else {
		fmt.Println("Unknown err")
		body, err := ioutil.ReadAll(resp.Body)
		util.CheckErr(err)

		return errors.New(string(body))
	}

	return nil
}
