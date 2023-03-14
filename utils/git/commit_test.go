package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitGetHeadCommit(t *testing.T) {
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
	commitId, err := GitGetHeadCommit(absPath)
	require.NoError(t, err)

	//assert
	require.Equal(t, "69115c87e97a21941d48db4dc04c3d6cc8380d0a", commitId)
}
