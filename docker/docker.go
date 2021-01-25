package docker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/stshontikidis/image_check/util"
)

func auth(repo string, scope string) string {
	service := "registry.docker.io"
	url := "https://auth.docker.io/token?service=" + service + "&scope=" + scope

	resp, err := http.Get(url)
	util.CheckErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var auth authResponse

	err = json.Unmarshal(body, &auth)
	if err != nil {
		panic(err)
	}

	return auth.Token
}

// GetDigest retrieves docker manifest digest
func GetDigest(repo string, tag string) string {
	scope := "repository:" + repo + ":pull"
	token := auth(repo, scope)

	url := "https://registry.hub.docker.com/v2/" + repo + "/manifests/stable"

	req, err := http.NewRequest("GET", url, nil)
	util.CheckErr(err)

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")

	client := &http.Client{}

	resp, err := client.Do(req)
	util.CheckErr(err)

	body, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err)

	var manifest manifest

	err = json.Unmarshal(body, &manifest)
	util.CheckErr(err)

	return manifest.Config.Digest
}
