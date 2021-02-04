package config

import (
	"os"
)

// Config general config structure
type Config struct {
	GithubRef        string
	GithubRepo       string
	GithubToken      string
	GithubWorkflowID string
	DockerRepo       string
	DockerTag        string
}

// GetConfig returns pointer to config
func GetConfig() *Config {
	githubRef := os.Getenv("GH_REF")
	githubRepo := os.Getenv("GH_REPO")
	token := os.Getenv("GH_TOKEN")
	dockerRepo := os.Getenv("DOCKER_REPO")
	dockerTag := os.Getenv("DOCKER_TAG")

	workflowID := os.Getenv("GH_WORKFLOW")

	cfg := &Config{
		GithubRef:        githubRef,
		GithubRepo:       githubRepo,
		GithubToken:      token,
		GithubWorkflowID: workflowID,
		DockerRepo:       dockerRepo,
		DockerTag:        dockerTag,
	}

	return cfg
}
