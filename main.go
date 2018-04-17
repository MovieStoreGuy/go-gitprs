package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MovieStoreGuy/prcheck/service"
	"github.com/fatih/color"
)

var (
	user, token string
	org, team   string
)

func init() {
	const (
		blank = ""
	)
	flag.StringVar(&user, "user", blank, "defines the user to examine")
	flag.StringVar(&token, "token", blank, "the github token to be used with the client")
	flag.StringVar(&org, "org", blank, "the organisation to examine")
	flag.StringVar(&team, "team", blank, "the team to examine")
}

func main() {
	flag.Parse()
	projects, err := service.New(user, org, team, token).GetOpenPrs()
	if err != nil {
		fmt.Println("Unable to process prs due to:", err)
		flag.Usage()
		os.Exit(-1)
	}
	for _, project := range projects {
		color.Green("Project name %v", project.Name)
		for _, pr := range project.PullRequests {
			color.Yellow("\t%v", pr.Title)
			color.Blue("\t%v", pr.Link)
		}
	}
}
