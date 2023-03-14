package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitGetRemote(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	remoteName := "origin"
	foundGitRemote, err := GitGetRemote(absPath, remoteName)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, gitRemote, foundGitRemote)
}

func TestGitSetRemote(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	customRemoteName := "custom2"
	customGitRemote := "https://github.com/AnalogJ/npm_analogj_test2.git"
	_, err = GitSetRemote(absPath, customRemoteName, customGitRemote)
	foundGitRemote, err := GitGetRemote(absPath, customRemoteName)
	require.NoError(t, err)

	//assert
	require.NoError(t, err)
	require.Equal(t, customGitRemote, foundGitRemote)
}
