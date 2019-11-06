package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"log"
)

type GithubClient struct {
	ctx    context.Context
	client *github.Client
}

type GithubIssue struct {
	issue *github.Issue
}

func NewClient(token string) *GithubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &GithubClient{
		ctx:    ctx,
		client: client,
	}
}

func (ghc *GithubClient) AddLabel(githubOrg string, githubRepo string, issueId int, label string) {
	_, _, err := ghc.client.Issues.AddLabelsToIssue(ghc.ctx, githubOrg, githubRepo, issueId, []string{label})
	if err != nil {
		log.Fatalln(err)
	}
}

func (ghc *GithubClient) FetchIssues(githubOrg, githubRepo, label string) (features []*GithubIssue, fixes []*GithubIssue, docs []*GithubIssue, pullRequests []*GithubIssue) {
	options := &github.IssueListByRepoOptions{State: "all", Labels: []string{label}}

	for {
		issues, response, err := ghc.client.Issues.ListByRepo(ghc.ctx, githubOrg, githubRepo, options)
		if err != nil {
			log.Fatalln(err)
		}

		for _, issue := range issues {
			ghi := &GithubIssue{issue: issue}
			if issue.IsPullRequest() {
				pullRequests = append(pullRequests, ghi)
			} else {
				for _, label := range issue.Labels {
					if *label.Name == "Type: Enhancement" {
						features = append(features, ghi)
					}
					if *label.Name == "Type: Bug" {
						fixes = append(fixes, ghi)
					}
					if *label.Name == "Type: Docs" {
						docs = append(docs, ghi)
					}
				}
			}
		}

		if response.NextPage == 0 {
			break
		}

		options.Page = response.NextPage
	}

	return
}

func PrintIssues(title string, issues []*GithubIssue) {
	if len(issues) > 0 {
		fmt.Println("##", title)
		for _, issue := range issues {
			fmt.Printf("* %s ([#%d](%s))\n", *issue.issue.Title, *issue.issue.Number, *issue.issue.HTMLURL)
		}
		fmt.Println()
	}
}
