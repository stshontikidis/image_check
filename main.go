package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stshontikidis/image_check/docker"
	"github.com/stshontikidis/image_check/github"
	"github.com/stshontikidis/image_check/util"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cfg := util.GetConfig()

	file, err := os.OpenFile("/tmp/digest", os.O_CREATE|os.O_RDWR, 0664)
	util.CheckErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	newDigest := docker.GetDigest(cfg.DockerRepo, "stable")
	oldDigest := scanner.Text()

	if oldDigest != newDigest {
		fmt.Println("no match!!")
		err := github.Dispatch(cfg.GithubRepo, cfg.GithubToken)
		util.CheckErr(err)

		util.WipeAndWrite(file, newDigest)
	}

}
