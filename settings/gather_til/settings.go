package gather_til

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type NotificationSetting interface {
	Post()
}

type Notification struct {
}

type Discord struct {
	WebhookUrl string
	Notification
}

func (d Discord) Post() {
}

type Slack struct {
	Token string
	Notification
}

func (d Slack) Post() {
}

type GitHubSettings struct {
	GitHubToken string
	RepoUserName string
	RepoName string
	TargetLabel string
}

type EnvSettings struct {
	GitHub               GitHubSettings
	NotificationSettings []NotificationSetting
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

func getNotificationSetting() []NotificationSetting {
	var result []NotificationSetting
	if key := os.Getenv("DISCORD_URL"); len(key) != 0 {
		discord := Discord{
			WebhookUrl: os.Getenv("DISCORD_URL"),
		}
		result = append(result, discord)
	}
	if key := os.Getenv("SLACK_TOKEN"); len(key) != 0 {
		slack := Slack{
			Token: os.Getenv("SLACK_TOKEN"),
		}
		result = append(result, slack)
	}
	return result
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

	notification := getNotificationSetting()

	return EnvSettings {
		GitHub: GitHubSettings{
			GitHubToken: os.Getenv("GITHUB_TOKEN"),
			RepoUserName: getDefaultEnv("REPO_USER_NAME"),
			RepoName:  getDefaultEnv("REPO_NAME"),
			TargetLabel:  getDefaultEnv("TARGET_LABEL"),
		},
		NotificationSettings: notification,
	}
}