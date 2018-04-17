package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/MovieStoreGuy/prcheck/types"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Github is an object to cache important information
type Github struct {
	user         string
	team         string
	token        string
	organisation string
	client       *github.Client
}

// New will configure a github ready for use
func New(user, org, team, token string) *Github {
	var authClient *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		authClient = oauth2.NewClient(context.Background(), ts)
	}
	return &Github{
		user:         user,
		team:         team,
		token:        token,
		organisation: org,
		client:       github.NewClient(authClient),
	}
}

func (g *Github) GetOpenPrs() ([]types.Project, error) {
	var (
		repos []*github.Repository
		err   error
	)
	switch {
	case g.organisation != "" && g.token != "":
		// Get the organisation open projects
		repos, err = g.getOrgProjects()
	case g.user != "" || g.token != "":
		// Get the users project
		repos, err = g.getUsersProjects()
	default:
		return nil, errors.New("Misconfigured github client")
	}
	if err != nil {
		return nil, err
	}
	return g.getAllOpenPrs(repos)
}

func (g *Github) getOrgProjects() ([]*github.Repository, error) {
	projects := []*github.Repository{}
	opt := &github.RepositoryListByOrgOptions{}
	for {
		items, resp, err := g.client.Repositories.ListByOrg(context.Background(), g.organisation, opt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, items...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return projects, nil
}

func (g *Github) getUsersProjects() ([]*github.Repository, error) {
	projects := []*github.Repository{}
	opt := &github.RepositoryListOptions{}
	for {
		items, resp, err := g.client.Repositories.List(context.Background(), g.user, opt)
		if err != nil {
			return nil, err
		}
		// TODO(Sean marciniak): Enable to filter by team
		projects = append(projects, items...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return projects, nil
}

func (g *Github) getAllOpenPrs(projects []*github.Repository) ([]types.Project, error) {
	collection := []types.Project{}
	for _, project := range projects {
		opt := &github.PullRequestListOptions{}
		for {
			items, resp, err := g.client.PullRequests.List(context.Background(), project.GetOwner().GetLogin(), project.GetName(), opt)
			if err != nil {
				return nil, err
			}
			// If there is no PRs to review, don't bother listing it
			if len(items) == 0 {
				break
			}
			// Do the processing here plz
			repo := types.Project{
				Name: project.GetName(),
			}
			for _, item := range items {
				repo.PullRequests = append(repo.PullRequests, types.PullRequest{
					Title: item.GetTitle(),
					Link:  item.GetHTMLURL(),
				})
			}
			collection = append(collection, repo)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}
	return collection, nil
}
