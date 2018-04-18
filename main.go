package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MovieStoreGuy/go-gitprs/service"
	"github.com/fatih/color"
)

var (
	token     string
	org, team string
)

func init() {
	const (
		blank = ""
	)
	flag.StringVar(&token, "token", blank, "the github token to be used with the client")
	flag.StringVar(&org, "org", blank, "the organisation to examine")
	flag.StringVar(&team, "team", blank, "the team to examine")
}

func main() {
	flag.Parse()
	projects, err := service.New(org, team, token).GetOpenPrs()
	if err != nil {
		fmt.Println("Unable to process prs due to:", err)
		flag.Usage()
		os.Exit(-1)
	}
	for project := range projects {
		if project.Error != nil {
			color.Red("Issue with getting Project information due to: %v", project.Error)
			continue
		}
		color.Green("Project: %v", project.Name)
		for _, pr := range project.PullRequests {
			color.Yellow("\t%v", pr.Title)
			color.Blue("\t%v", pr.Link)
		}
	}
}
