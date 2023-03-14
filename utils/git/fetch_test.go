package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitFetchPullRequest(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	pullRequestNumber := "16"
	localBranchName := "test_fetch_branch"
	srcPatternTmpl := "refs/pull/%s/merge"
	destPatternTmpl := "refs/remotes/origin/pr/%s/merge"
	err = GitFetchPullRequest(absPath, pullRequestNumber, localBranchName, srcPatternTmpl, destPatternTmpl)
	require.NoError(t, err)
	commitId, err := GitGetHeadCommit(absPath)
	branch, err := GitGetBranch(absPath)

	//assert
	require.NoError(t, err)

	require.Equal(t, "acccf80ecd287afd6d5cb6975d3114c52fe4e03e", commitId) //merge commit for PR 16
	require.Equal(t, localBranchName, branch)
}
