package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stshontikidis/image_check/config"
	"github.com/stshontikidis/image_check/docker"
	"github.com/stshontikidis/image_check/github"
)

func main() {
	cfg := config.GetConfig()

	file, err := os.OpenFile("/tmp/digest", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		log.Fatal(err)
	}

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

func loop(digest *string, file *os.File, cfg *config.Config) {
	newDigest, err := docker.GetDigest(cfg.DockerRepo, cfg.DockerTag)
	if err != nil {
		log.Fatal(err)
	}

	if *digest != newDigest {
		fmt.Println("no match!!")
		err := github.Dispatch(cfg.GithubRepo,
			cfg.GithubRef,
			cfg.GithubWorkflowID,
			cfg.GithubToken,
			newDigest,
		)
		if err != nil {
			log.Fatal(err)
		}

		wipeAndWrite(file, newDigest)
		*digest = newDigest
	}
}

func wipeAndWrite(f *os.File, contents string) {
	f.Truncate(0)
	f.Seek(0, 0)
	f.WriteString(contents)
}
