package scm

import (
	"fmt"
	"github.com/packagrio/go-common/config"
	"github.com/packagrio/go-common/pipeline"
	"github.com/packagrio/go-common/utils/git"
	"net/http"
)

type scmBase struct {
	PipelineData *pipeline.Data
}

// configure method will generate an authenticated client that can be used to comunicate with Github
// MUST set @git_parent_path
// MUST set @client field
func (s *scmBase) Init(pipelineData *pipeline.Data, myConfig config.BaseInterface, httpClient *http.Client) error {
	s.PipelineData = pipelineData

	return nil
}

//We cant make any assumptions about the SCM or the environment. (No Pull requests or SCM env vars). So lets use native git methods to get
// the current repo status.
func (g *scmBase) RetrievePayload() (*Payload, error) {

	g.PipelineData.IsPullRequest = false

	//check the local git repo for relevant info
	remoteUrl, err := git.GitGetRemote(g.PipelineData.GitLocalPath, "origin")
	if err != nil {
		return nil, err
	}

	commit, err := git.GitGetHeadCommit(g.PipelineData.GitLocalPath)
	if err != nil {
		return nil, err
	}

	branch, err := git.GitGetBranch(g.PipelineData.GitLocalPath)
	if err != nil {
		return nil, err
	}

	return &Payload{
		Head: &pipeline.ScmCommitInfo{
			Sha: commit,
			Ref: branch,
			Repo: &pipeline.ScmRepoInfo{
				CloneUrl: remoteUrl,
			}},
	}, nil
}

func (s *scmBase) Publish() error {
	perr := git.GitPush(s.PipelineData.GitLocalPath, s.PipelineData.GitLocalBranch, s.PipelineData.GitBaseInfo.Ref, fmt.Sprintf("v%s", s.PipelineData.ReleaseVersion))
	if perr != nil {
		return perr
	}

	// calculate the release sha
	releaseCommit, err := git.GitGetHeadCommit(s.PipelineData.GitLocalPath)
	if err != nil {
		return err
	}
	s.PipelineData.ReleaseCommit = releaseCommit

	return nil
}
