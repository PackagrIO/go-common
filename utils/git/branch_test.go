package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitGetBranch(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)
	err = GitCheckout(absPath, "do_not_delete_capsulecd_test_branch")
	require.NoError(t, err)

	//test
	branch, err := GitGetBranch(absPath)

	//assert
	require.NoError(t, err)
	require.Equal(t, "do_not_delete_capsulecd_test_branch", branch)
}

func TestGitCreateBranchFromHead(t *testing.T) {

	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)
	err = GitCheckout(absPath, "do_not_delete_capsulecd_test_branch")
	require.NoError(t, err)

	//test
	branch, err := GitCreateBranchFromHead(absPath, "test_branch")

	//assert
	require.NoError(t, err)
	require.Equal(t, "test_branch", branch)
}
