package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Plugin defines the Downstream plugin parameters.
type Plugin struct {
	GithubToken string
	Number      int
	Action      string
	Message     string
	RepoOwner   string
	RepoName    string
}

// Exec runs the plugin
func (p *Plugin) Exec() error {
	if len(p.GithubToken) == 0 {
		return fmt.Errorf("Error: you must provide your Github access token.")
	}

	// Instantiate Github client
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: p.GithubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	// This could be a case
	// Switching for actions :
	if p.Action == "close" {
		return close(client, p)
	} else if p.Action == "comment" {
		return comment(client, p)
	} else if p.Action == method.Merge || p.Action == method.Squash || p.Action == method.Rebase {
		return merge(client, p, p.Action)
	}
	return fmt.Errorf("Error: you must provide a valid action.")
}

// Close a PR
func close(client *github.Client, p *Plugin) error {
	// If a message is set, leave it before closing
	if len(p.Message) > 0 {
		err := comment(client, p)
		if err != nil {
			return err
		}
	}
	_, _, err := client.PullRequests.Edit(oauth2.NoContext, p.RepoOwner, p.RepoName, p.Number, &github.PullRequest{
		State: &state.Closed,
	})
	if err != nil {
		return err
	}
	return nil
}

// Comment a PR
func comment(client *github.Client, p *Plugin) error {
	_, _, err := client.Issues.CreateComment(
		oauth2.NoContext, p.RepoOwner, p.RepoName, p.Number,
		&github.IssueComment{
			Body: &p.Message,
		})
	if err != nil {
		return err
	}
	return nil
}

// Merge a PR
func merge(client *github.Client, p *Plugin, mergeMethod string) error {
	// If a message is set, leave it before merging
	if len(p.Message) > 0 {
		err := comment(client, p)
		if err != nil {
			return err
		}
	}
	_, _, err := client.PullRequests.Merge(oauth2.NoContext, p.RepoOwner, p.RepoName, p.Number, p.Message,
		&github.PullRequestOptions{
			MergeMethod: mergeMethod,
		})
	if err != nil {
		return err
	}
	return nil
}

var state = struct {
	Closed string
	Open   string
}{"closed", "open"}

var method = struct {
	Merge  string
	Squash string
	Rebase string
}{"merge", "squash", "rebase"}
