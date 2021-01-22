package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	type AuthResponse struct {
		Token     string `json:"token"`
		ExpiresIn int    `json:"expires_in"`
		IssuedAt  string `json:"issued_at"`
	}

	type Config struct {
		Digest    string `json:"digest"`
		MediaType string `json:"mediaType"`
		Size      int    `json:"size"`
	}

	type Layer struct {
		Digest    string `json:"digest"`
		MediaType string `json:"mediaType"`
		Size      int    `json:"Size"`
	}
	type Manifest struct {
		Config        Config  `json:"config"`
		Layers        []Layer `json:"layers"`
		MediaType     string  `json:"mediaType"`
		SchemaVersion int     `json:"schemaVersion"`
	}

	service := "registry.docker.io"
	repo := "library/nextcloud"
	scope := "repository:" + repo + ":pull"
	url := "https://auth.docker.io/token?service=" + service + "&scope=" + scope

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var auth AuthResponse

	err = json.Unmarshal(body, &auth)
	if err != nil {
		fmt.Println(err)
	}

	url = "https://registry.hub.docker.com/v2/" + repo + "/manifests/stable"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Authorization", "Bearer "+auth.Token)
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", resp)

	body, err = ioutil.ReadAll(resp.Body)

	sb := string(body)
	fmt.Println(sb)

	var manifest Manifest

	err = json.Unmarshal(body, &manifest)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", manifest)
}
