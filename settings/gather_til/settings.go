package gather_til

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)


type GitHubSettings struct {
	GitHubToken string
	RepoUserName string
	RepoName string
	TargetLabel string
}

type EnvSettings struct {
	Github GitHubSettings
}

var defaultSettingValues = map[string]string {
	"EXEC_MODE": "dev",
}

const AppName = "gather_til"

func getDefaultEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultSettingValues[key]
	}
	return value
}

func GetEnv() EnvSettings {
	mode := getDefaultEnv("EXEC_MODE")
	filename := strings.Join([]string{".env", AppName, mode}, ".")
	if _, err := os.Stat(filename); err != nil {
		fmt.Printf("Skip Load env file: filename=%s\n", filename)
	} else {
		err := godotenv.Load(filename)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			panic(err)
		} else {
			fmt.Printf("Load env file: filename=%s\n", filename)
		}
	}

	return EnvSettings {
		GitHubSettings{
			GitHubToken: os.Getenv("GITHUB_TOKEN"),
			RepoUserName: getDefaultEnv("REPO_USER_NAME"),
			RepoName:  getDefaultEnv("REPO_NAME"),
			TargetLabel:  getDefaultEnv("TARGET_LABEL"),
		},
	}
}