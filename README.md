# Go Git PRs

A Golang application to notify you if there is any outstanding PRs that require your attention


## Install
To install the tool, you'll need golang installed on your machine and you can simply
```sh
go get -u github.com/MovieStoreGuy/go-gitprs
```

## Usage
To fetch all PRs that you have opened simply:
```sh
go-gitprs --token ${GITHUB_TOKEN}
```
Where the Github token is a oauth token from github that has repo read access.

To fetch all PRs opened by an organisation you belong to, simply:
```sh
go-gitprs --token ${GITHUB_TOKEN} --org <OrgName>
```

This will fetch all open PRs that you could review within your github organisation.

Then to filter based off a given team, simply:
```sh
go-gitprs --token ${GITHUB_TOKEN} --org <OrgName> --team <TeamName>
```
