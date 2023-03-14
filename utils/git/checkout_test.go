package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestGitCheckout(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"

	//test
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	err = GitCheckout(absPath, "do_not_delete_capsulecd_test_branch")
	require.NoError(t, err)

	//assert
	require.Equal(t, absPath, path.Join(parentPath, repoName))
	require.FileExists(t, path.Join(absPath, "branch_specific_file.txt"), "branch_specific_file.txt should exist in the cloned repo custom branch")
}
