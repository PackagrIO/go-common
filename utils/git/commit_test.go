package git

import (
	"github.com/stretchr/testify/require"
	git "gopkg.in/libgit2/git2go.v25"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
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

func TestGitCommit(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	//using a constant signature so that the commit hash is always the same.
	err = ioutil.WriteFile(path.Join(absPath, "test.txt"), []byte("test"), 0644)
	require.NoError(t, err)
	sig := git.Signature{Name: "test", Email: "test@example.com", When: time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)}
	err = GitCommit(absPath, "test commit", &sig)
	require.NoError(t, err)
	headCommit, err := GitGetHeadCommit(absPath)
	require.NoError(t, err)

	//assert
	require.Equal(t, "816eca4d20a2a27664d8681ad5d2b89757dcd0bc", headCommit)
}
