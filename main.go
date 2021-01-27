package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/stshontikidis/image_check/docker"
	"github.com/stshontikidis/image_check/github"
	"github.com/stshontikidis/image_check/util"
)

func loop(digest *string, file *os.File, cfg *util.Config) {
	newDigest := docker.GetDigest(cfg.DockerRepo, "stable")

	if *digest != newDigest {
		fmt.Println("no match!!")
		err := github.Dispatch(cfg.GithubRepo,
			cfg.GithubRef,
			cfg.GithubWorkflowID,
			cfg.GithubToken,
			newDigest,
		)
		util.CheckErr(err)

		util.WipeAndWrite(file, newDigest)
		*digest = newDigest
	}
}

func main() {
	cfg := util.GetConfig()

	file, err := os.OpenFile("/tmp/digest", os.O_CREATE|os.O_RDWR, 0664)
	util.CheckErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	digest := scanner.Text()
	for {
		fmt.Println("looping")
		loop(&digest, file, cfg)
		time.Sleep(time.Hour)
	}
}
