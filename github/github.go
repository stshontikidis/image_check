package github

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/stshontikidis/image_check/util"
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

	util.CheckErr(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	util.CheckErr(err)

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	util.CheckErr(err)

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, err := ioutil.ReadAll(resp.Body)
		util.CheckErr(err)
		return errors.New(string(body))
	}

	return nil
}
