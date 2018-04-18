package types

type Project struct {
	Name         string
	PullRequests []PullRequest
	Error        error
}

type PullRequest struct {
	Title string
	Link  string
}
