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
			Usage:  "Github auth token",
			EnvVar: "GITHUB_TOKEN,GITHUB_PR_GITHUB_TOKEN,PLUGIN_GITHUB_TOKEN",
		},
		cli.IntFlag{
			Name:   "number",
			Usage:  "PR number to work with",
			EnvVar: "GITHUB_SEARCH_DOWNSTREAM_BRANCH,PLUGIN_NUMBER,DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "action",
			Usage:  "Trigger a drone build on a custom server",
			EnvVar: "DRONE_SERVER,GITHUB_SEARCH_DOWNSTREAM_DRONE_SERVER,PLUGIN_DRONE_SERVER",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "Drone API token from your user settings",
			EnvVar: "DRONE_TOKEN,GITHUB_SEARCH_DOWNSTREAM_DRONE_TOKEN,PLUGIN_DRONE_TOKEN",
		},
		cli.StringFlag{
			Name:   "repo-owner",
			Usage:  "Repo owner",
			EnvVar: "DRONE_REPO_OWNER",
		},

		cli.StringFlag{
			Name:   "repo-name",
			Usage:  "Repo name",
			EnvVar: "DRONE_REPO_NAME",
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
