package main

import (
	"fmt"
	"github.com/yumechi/gather_TIL/settings/gather_til"
	"reflect"
)

type Discord struct {
	gather_til.Discord
}

func (d Discord) Post() {
	fmt.Printf("This is Discord: WebhookUrl=%s\n", d.WebhookUrl)
}


func main() {
	settings := gather_til.GetEnv()

	// 通知系を送り付ける
	for _, elem := range settings.NotificationSettings {
		elemName := reflect.TypeOf(elem).Name()
		println(elemName)

		switch elem.(type) {
		case gather_til.Discord:
			d := Discord{
				elem.(gather_til.Discord),
			}
			d.Post()
		case gather_til.Slack:
			fmt.Println("This is Slack")
		default:
			fmt.Printf("Unknown Type: %s\n", elemName)
		}
	}
}