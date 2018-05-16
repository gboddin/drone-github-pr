package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "github PR plugin"
	app.Usage = "github PR plugin"
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "github-token",
			Usage:  "Github OAuth token",
			EnvVar: "GITHUB_TOKEN,GITHUB_PR_GITHUB_TOKEN,PLUGIN_GITHUB_TOKEN",
		},
		cli.IntFlag{
			Name:   "number",
			Usage:  "PR number",
			EnvVar: "GITHUB_PR_NUMBER,PLUGIN_NUMBER,DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "action",
			Usage:  "Action to trigger ( one of comment, close, merge, rebase, squash )",
			EnvVar: "GITHUB_PR_ACTION,PLUGIN_ACTION",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "Comment to leave on PR and in merge commit",
			EnvVar: "GITHUB_PR_MESSAGE,PLUGIN_MESSAGE,DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "repo-owner",
			Usage:  "Repo owner",
			EnvVar: "GITHUB_PR_REPO_OWNER,PLUGIN_REPO_OWNER,DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo-name",
			Usage:  "Repo name",
			EnvVar: "GITHUB_PR_REPO_NAME,PLUGIN_REPO_NAME,DRONE_REPO_NAME",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		GithubToken: c.String("github-token"),
		Message:     c.String("message"),
		Number:      c.Int("number"),
		Action:      c.String("action"),
		RepoName:    c.String("repo-name"),
		RepoOwner:   c.String("repo-owner"),
	}
	return plugin.Exec()
}
