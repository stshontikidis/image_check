package docker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func auth(repo string, scope string) (string, error) {
	service := "registry.docker.io"
	url := "https://auth.docker.io/token?service=" + service + "&scope=" + scope
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "", fmt.Errorf("Auth timed out to %s", url)
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf(
			"Authentication failed at %s with service %s and scope %s response %s", url, service, scope, b,
		)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var auth authResponse

	err = json.Unmarshal(body, &auth)
	if err != nil {
		return "", fmt.Errorf("Unable to parse auth response %s with error %s", body, err)
	}

	return auth.Token, nil
}

// GetDigest retrieves docker manifest digest
func GetDigest(repo string, tag string) (string, error) {
	scope := "repository:" + repo + ":pull"
	token, err := auth(repo, scope)
	if err != nil {
		return "", err
	}

	url := "https://registry.hub.docker.com/v2/" + repo + "/manifests/" + tag

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		err, ok := err.(net.Error)
		if ok && err.Timeout() {
			return "", fmt.Errorf("Manifest request timed out %s", url)
		}
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("Failed to get image manifest from %s response %s", req.URL, b)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var manifest manifest

	err = json.Unmarshal(body, &manifest)
	if err != nil {
		return "", fmt.Errorf("Unable to parse manifest respone %s with error %s", body, err)
	}

	if manifest.Config.Digest == "" {
		return "", fmt.Errorf("No digest found in manifest %+v", manifest)
	}

	return manifest.Config.Digest, nil
}
