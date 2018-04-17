package types

type Project struct {
	Name         string
	PullRequests []PullRequest
}

type PullRequest struct {
	Title string
	Link  string
}
