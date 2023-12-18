package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Config struct {
	GitHubLoginConfig oauth2.Config
}

var AppConfig Config

func GithubConfig() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.GitHubLoginConfig = oauth2.Config{
		RedirectURL: "http://localhost:8080/github_callback",
		ClientID:    os.Getenv("GITHUB_CLIENT_ID"),
		//RedirectURL: fmt.Sprintf(
		//	"https://github.com/login/oauth/authorize?scope=user:repo&client_id=%s&redirect_uri=%s", os.Getenv("GITHUB_CLIENT_ID"), "http://localhost:8080/github_callback"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}

	return AppConfig.GitHubLoginConfig
}
