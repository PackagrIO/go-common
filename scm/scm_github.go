package scm

import (
	"encoding/json"
	"fmt"
	"github.com/analogj/go-util/utils"
	"github.com/google/go-github/v32/github"
	"github.com/packagrio/go-common/errors"
	"github.com/packagrio/go-common/pipeline"
	"io/ioutil"
	"os"
	"strconv"
)

// TAKEN from: https://github.com/google/go-github/blob/master/github/event_types.go
// TODO: this is not yet available in master, once it is, we should remove this Struct.
//
// WorkflowDispatchEvent is triggered when someone triggers a workflow run on GitHub or
// sends a POST request to the create a workflow dispatch event endpoint.
//
// GitHub API docs: https://docs.github.com/en/developers/webhooks-and-events/webhook-events-and-payloads#workflow_dispatch
type WorkflowDispatchEvent struct {
	Inputs   json.RawMessage `json:"inputs,omitempty"`
	Ref      *string         `json:"ref,omitempty"`
	Workflow *string         `json:"workflow,omitempty"`

	// The following fields are only populated by Webhook events.
	Repo   *github.Repository   `json:"repository,omitempty"`
	Org    *github.Organization `json:"organization,omitempty"`
	Sender *github.User         `json:"sender,omitempty"`
}

type scmGithub struct {
	PipelineData *pipeline.Data
}

func (g *scmGithub) Init(pipelineData *pipeline.Data) error {
	g.PipelineData = pipelineData

	//if _, present := os.LookupEnv("GITHUB_TOKEN"); !present {
	//	return errors.ScmAuthenticationFailed("Missing github access token")
	//}
	//if g.Config.IsSet("scm_git_parent_path") {
	//	g.PipelineData.GitParentPath = g.Config.GetString("scm_git_parent_path")
	//	os.MkdirAll(g.PipelineData.GitParentPath, os.ModePerm)
	//} else {
	//	dirPath, _ := ioutil.TempDir("", "")
	//	g.PipelineData.GitParentPath = dirPath
	//}

	//if _, present := os.LookupEnv("GITHUB_ACTION"); !present {
	//	//running as a github action.
	//
	//} else {
	//	ctx := context.Background()
	//	ts := oauth2.StaticTokenSource(
	//		&oauth2.Token{AccessToken: g.Config.GetString("scm_github_access_token")},
	//	)
	//	tc := oauth2.NewClient(ctx, ts)
	//
	//	//TODO: autopaginate turned on.
	//	g.Client = github.NewClient(tc)
	//}

	//if client != nil {
	//	//primarily used for testing.
	//	g.Client = github.NewClient(client)
	//} else {
	//	ctx := context.Background()
	//	ts := oauth2.StaticTokenSource(
	//		&oauth2.Token{AccessToken: g.Config.GetString("scm_github_access_token")},
	//	)
	//	tc := oauth2.NewClient(ctx, ts)
	//
	//	//TODO: autopaginate turned on.
	//	g.Client = github.NewClient(tc)
	//}
	//
	//if g.Config.IsSet("scm_github_api_endpoint") {
	//
	//	apiUrl, aerr := url.Parse(g.Config.GetString("scm_github_api_endpoint"))
	//	if aerr != nil {
	//		return aerr
	//	}
	//	g.Client.BaseURL = apiUrl
	//}

	return nil
}

func (g *scmGithub) RetrievePayload() (*Payload, error) {
	eventType := utils.GetEnv("GITHUB_EVENT_NAME", "push")
	eventPayloadPath, present := os.LookupEnv("GITHUB_EVENT_PATH")
	if !present {
		return nil, errors.ScmPayloadFormatError("Event Payload not present")
	}

	//open & parse JSON File
	jsonBytes, err := ioutil.ReadFile(eventPayloadPath)
	if err != nil {
		return nil, errors.ScmFilesystemError("Event Payload Path does not exist")
	}

	if eventType == "push" {
		var pushEvent github.PushEvent
		err := json.Unmarshal(jsonBytes, &pushEvent)
		if err != nil {
			return nil, errors.ScmPayloadFormatError(err.Error())
		}

		g.PipelineData.IsPullRequest = false

		return &Payload{
			Head: &pipeline.ScmCommitInfo{
				Sha: pushEvent.GetAfter(),
				Ref: pushEvent.GetRef(),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: pushEvent.GetRepo().GetCloneURL(),
					Name:     pushEvent.GetRepo().GetName(),
					FullName: pushEvent.GetRepo().GetFullName(),
				}},
		}, nil
		//make this as similar to a pull request as possible
	} else if eventType == "pull_request" {

		//parse Pull Request event payload
		var pullRequestEvent github.PullRequestEvent
		err := json.Unmarshal(jsonBytes, &pullRequestEvent)
		if err != nil {
			return nil, errors.ScmPayloadFormatError(err.Error())
		}

		g.PipelineData.IsPullRequest = true

		return &Payload{
			Title:             pullRequestEvent.GetPullRequest().GetTitle(),
			PullRequestNumber: strconv.Itoa(pullRequestEvent.GetPullRequest().GetNumber()),
			Head: &pipeline.ScmCommitInfo{
				Sha: pullRequestEvent.GetPullRequest().GetHead().GetSHA(),
				Ref: pullRequestEvent.GetPullRequest().GetHead().GetRef(),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: pullRequestEvent.GetPullRequest().GetHead().GetRepo().GetCloneURL(),
					Name:     pullRequestEvent.GetPullRequest().GetHead().GetRepo().GetName(),
					FullName: pullRequestEvent.GetPullRequest().GetHead().GetRepo().GetFullName(),
				},
			},
			Base: &pipeline.ScmCommitInfo{
				Sha: pullRequestEvent.GetPullRequest().GetBase().GetSHA(),
				Ref: pullRequestEvent.GetPullRequest().GetBase().GetRef(),
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: pullRequestEvent.GetPullRequest().GetBase().GetRepo().GetCloneURL(),
					Name:     pullRequestEvent.GetPullRequest().GetBase().GetRepo().GetName(),
					FullName: pullRequestEvent.GetPullRequest().GetBase().GetRepo().GetFullName(),
				},
			},
		}, nil
	} else if eventType == "workflow_dispatch" {
		//parse Workflow Dispatch (manual) event payload
		var wfDispatchEvent WorkflowDispatchEvent
		err := json.Unmarshal(jsonBytes, &wfDispatchEvent)
		if err != nil {
			return nil, errors.ScmPayloadFormatError(err.Error())
		}

		g.PipelineData.IsPullRequest = false
		return &Payload{
			Head: &pipeline.ScmCommitInfo{
				//Sha: wfDispatchEvent.GetAfter(),
				Ref: *wfDispatchEvent.Ref,
				Repo: &pipeline.ScmRepoInfo{
					CloneUrl: wfDispatchEvent.Repo.GetCloneURL(),
					Name:     wfDispatchEvent.Repo.GetName(),
					FullName: wfDispatchEvent.Repo.GetFullName(),
				}},
		}, nil

	} else {
		return nil, errors.ScmPayloadUnsupported("Unknown Event Type. Exiting.")
	}
}

// see https://github.com/actions/toolkit/blob/main/docs/commands.md
func (g *scmGithub) SetEnvironmentalVariable(name string, value string) error {
	fmt.Printf("\n::set-env name=%s::%s\n", name, value)
	return nil
}

// To prepend a string to PATH
func (g *scmGithub) AddPath(path string) error {
	fmt.Printf("\n::add-path::%s\n", path)
	return nil
}

// To set an output for the step
func (g *scmGithub) SetOutput(name string, value string) error {
	fmt.Printf("\n::set-output name=%s::%s\n", name, value)
	return nil
}

// To mask a value in the logs
func (g *scmGithub) MaskSecret(secret string) error {
	fmt.Printf("\n::add-mask::%s\n", secret)
	return nil
}
