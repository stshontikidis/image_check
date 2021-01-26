package util

import (
	"os"
)

// Config general config structure
type Config struct {
	GithubToken string
	Repo        string
}

// GetConfig returns pointer to config
func GetConfig() *Config {
	token := os.Getenv("GH_TOKEN")
	repo := os.Getenv("REPO")

	cfg := &Config{
		GithubToken: token,
		Repo:        repo,
	}

	return cfg
}
