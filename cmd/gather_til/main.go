package main

import "github.com/yumechi/gather_TIL/settings/gather_til"

func main() {
	settings := gather_til.GetEnv()
	// debug
	println(settings.Github.GitHubToken)
}