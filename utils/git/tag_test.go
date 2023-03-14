package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitTag(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	signature := GitSignature("test", "test@example.com")
	tagCommit, err := GitTag(absPath, "test_tag", "test_tag_message", signature)
	require.NoError(t, err)
	foundTagCommmit, err := GitGetTagDetails(absPath, "test_tag")
	require.NoError(t, err)

	//assert
	require.Equal(t, foundTagCommmit.CommitSha, tagCommit)
}

func TestGitGetTagDetails_Lightweight(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	tagName := "v1.0.10"
	tagDetails, err := GitGetTagDetails(absPath, tagName)
	require.NoError(t, err)

	//assert
	require.Equal(t, tagName, tagDetails.TagShortName)
	require.Equal(t, "86c62a8bf3c269eb6f82e9f8f1bc0646d5f46ef1", tagDetails.CommitSha)
}

func TestGitGetTagDetails_Annotated(t *testing.T) {
	//setup
	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
	parentPath := os.TempDir()
	repoName := "test_repo"
	absPath, err := GitClone(parentPath, repoName, gitRemote)
	require.NoError(t, err)
	defer os.RemoveAll(absPath)

	//test
	tagName := "v1.0.11"
	tagDetails, err := GitGetTagDetails(absPath, tagName)
	require.NoError(t, err)

	//assert
	require.Equal(t, tagName, tagDetails.TagShortName)
	require.Equal(t, "971b29d2f4d2d3797fe5b44488e12d1865ff71d0", tagDetails.CommitSha)
}

//TODO: enable this
//func TestGitFindNearestTagName(t *testing.T) {
//	//setup
//	gitRemote := "https://github.com/AnalogJ/npm_analogj_test.git"
//	parentPath := os.TempDir()
//	repoName := "test_repo"
//	absPath, err := GitClone(parentPath, repoName, gitRemote)
//	require.NoError(t, err)
//	defer os.RemoveAll(absPath)
//	err = GitCheckout(absPath, "nearest_tag")
//	require.NoError(t, err)
//
//	//test
//	nearestTag, err := GitFindNearestTagName(absPath)
//	require.NoError(t, err)
//
//	//assert
//	require.Equal(t, "v1.0.9", nearestTag)
//}
