package github

import (
	"github.com/google/go-github/v32/github"
	"github.com/packagrio/go-common/pipeline"
	"github.com/packagrio/go-common/scm"
	"strconv"
)

func PayloadFromGithubPullRequest(pullRequest github.PullRequest) *scm.Payload {
	return &scm.Payload{
		Title:             pullRequest.GetTitle(),
		PullRequestNumber: strconv.Itoa(pullRequest.GetNumber()),
		Head: &pipeline.ScmCommitInfo{
			Sha: pullRequest.GetHead().GetSHA(),
			Ref: pullRequest.GetHead().GetRef(),
			Repo: &pipeline.ScmRepoInfo{
				CloneUrl: pullRequest.GetHead().GetRepo().GetCloneURL(),
				Name:     pullRequest.GetHead().GetRepo().GetName(),
				FullName: pullRequest.GetHead().GetRepo().GetFullName(),
			},
		},
		Base: &pipeline.ScmCommitInfo{
			Sha: pullRequest.GetBase().GetSHA(),
			Ref: pullRequest.GetBase().GetRef(),
			Repo: &pipeline.ScmRepoInfo{
				CloneUrl: pullRequest.GetBase().GetRepo().GetCloneURL(),
				Name:     pullRequest.GetBase().GetRepo().GetName(),
				FullName: pullRequest.GetBase().GetRepo().GetFullName(),
			},
		},
	}
}
