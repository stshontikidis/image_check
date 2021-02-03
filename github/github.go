package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type input struct {
	Digest string `json:"digest"`
}

type payload struct {
	Ref    string `json:"ref"`
	Inputs input  `json:"inputs"`
}

// Dispatch will execute a repository_dispatch event against github repo
func Dispatch(repo string, ref string, workflowID string, token string, digest string) error {
	path := "repos/" + repo + "/actions/workflows/" + workflowID + "/dispatches"
	url := "https://api.github.com/" + path

	body, err := json.Marshal(
		payload{Ref: ref,
			Inputs: input{
				Digest: digest,
			},
		})

	if err != nil {
		return fmt.Errorf("Unable to format POST body")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		err, ok := err.(net.Error)
		if ok && err.Timeout() {
			return fmt.Errorf("POST request timed out %s", url)
		}
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}
