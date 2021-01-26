package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/stshontikidis/image_check/docker"
	"github.com/stshontikidis/image_check/util"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cfg := util.GetConfig()
	repo := cfg.Repo
	fmt.Println(repo)
	file, err := os.OpenFile("/tmp/digest", os.O_CREATE|os.O_RDWR, 0664)
	util.CheckErr(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	newDigest := docker.GetDigest(repo, "stable")
	oldDigest := scanner.Text()

	if oldDigest != newDigest {
		fmt.Println("no match!!")
		util.WipeAndWrite(file, newDigest)
	}

}
