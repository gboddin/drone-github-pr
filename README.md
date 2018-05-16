# drone-github-pr

[![Build Status](https://hold-on.nobody.run/api/badges/gboddin/drone-github-pr/status.svg)](http://hold-on.nobody.run/drone-github-pr)

Drone plugin to work on Github pull requests.

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/github-pr
docker build --rm -t gboo/github-pr .
```

## Usage

```
GLOBAL OPTIONS:
   --github-token value  Github OAuth token [$GITHUB_TOKEN, $GITHUB_PR_GITHUB_TOKEN, $PLUGIN_GITHUB_TOKEN]
   --number value        PR number (default: 0) [$GITHUB_PR_NUMBER, $PLUGIN_NUMBER, $DRONE_PULL_REQUEST]
   --action value        Action to trigger ( one of comment, close, merge, rebase, squash ) [$GITHUB_PR_ACTION, $PLUGIN_ACTION]
   --message value       Comment to leave on PR and in merge commit [$GITHUB_PR_MESSAGE, $PLUGIN_MESSAGE, $DRONE_COMMIT_MESSAGE]
   --repo-owner value    Repo owner [$GITHUB_PR_REPO_OWNER, $PLUGIN_REPO_OWNER, $DRONE_REPO_OWNER]
   --repo-name value     Repo name [$GITHUB_PR_REPO_NAME, $PLUGIN_REPO_NAME, $DRONE_REPO_NAME]
```


Execute from the working directory:

```sh
./drone-github-pr --github-token=aaeae7ae7ae7ae9ae897eaae97ae97a --action=comment --number=5 --repo-owner=Octocat --repo-name=drone-test --message="Hello world"

```

From Drone:

```yaml
pipeline:
  comment-pr:
    image: gboo/github-pr
    secrets: [ github_token ]
    action: comment
    message: "This PR looks good so far!"
    when:
      event: pull_request

  merge-pr:
    image: gboo/github-pr
    secrets: [ github_token ]
    action: rebase
    message: "Merging this PR"
    when:
      event: pull_request
```
