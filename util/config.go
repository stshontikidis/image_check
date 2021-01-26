package util

import (
	"os"
)

// Config general config structure
type Config struct {
	GithubRepo  string
	GithubToken string
	DockerRepo  string
}

// GetConfig returns pointer to config
func GetConfig() *Config {
	githubRepo := os.Getenv("GH_REPO")
	token := os.Getenv("GH_TOKEN")
	dockerRepo := os.Getenv("DOCKER_REPO")

	cfg := &Config{
		GithubRepo:  githubRepo,
		GithubToken: token,
		DockerRepo:  dockerRepo,
	}

	return cfg
}
