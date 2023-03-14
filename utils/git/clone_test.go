package git

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestGitClone(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"

	//test
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	defer os.RemoveAll(absPath)

	//assert
	require.NoError(t, err)
	require.Equal(t, absPath, path.Join(parentPath, repoName))
	require.FileExists(t, path.Join(absPath, "package.json"), "package.json should exist in the cloned repo")
}

func TestGitCloneWithGithubToken(t *testing.T) {
	if os.Getenv("GITHUB_TOKEN") == "" {
		t.Errorf("Skipping test because it requires a valid git remote with embedded credentials.")
	}

	//setup
	gitRemote := fmt.Sprintf("https://%s:%s@github.com/PackagrIO/test_npm_private_repo.git", "PackagrIO", os.Getenv("GITHUB_TOKEN"))
	parentPath := os.TempDir()
	repoName := "test_repo"

	//test
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	defer os.RemoveAll(absPath)

	//assert
	require.NoError(t, err)
	require.Equal(t, absPath, path.Join(parentPath, repoName))
	require.FileExists(t, path.Join(absPath, "package.json"), "package.json should exist in the cloned repo")
}
