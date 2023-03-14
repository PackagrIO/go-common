package git

// TODO: fix this
//func TestGitPush(t *testing.T) {
//	if os.Getenv("GITHUB_TOKEN") == "" {
//		t.Errorf("Skipping test because it requires a valid git remote with embedded credentials.")
//	}
//
//	//setup
//	gitRemote := fmt.Sprintf("https://%s:%s@github.com/PackagrIO/test_npm_private_repo.git", "PackagrIO", os.Getenv("GITHUB_TOKEN"))
//	parentPath := os.TempDir()
//	repoName := "test_repo"
//	absPath, err := GitClone(parentPath, repoName, gitRemote)
//	defer os.RemoveAll(absPath)
//
//	//checkout a PR
//	//create test file
//	//commit test file
//	//create test tag
//	localBranchName := "local_new_branch_push_pr_testing"
//	srcPatternTmpl := "refs/pull/%s/merge"
//	destPatternTmpl := "refs/remotes/origin/pr/%s/merge"
//	err = GitFetchPullRequest(absPath, "1", localBranchName, srcPatternTmpl, destPatternTmpl)
//	require.NoError(t, err)
//
//	err = ioutil.WriteFile(path.Join(absPath, "test.txt"), []byte("test"), 0644)
//	require.NoError(t, err)
//
//	sig := object.Signature{Name: "test", Email: "test@example.com", When: time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC)}
//	err = GitCommit(absPath, "test commit", &sig)
//	require.NoError(t, err)
//
//	testRandomSuffix := randomString(10)
//	tagName := fmt.Sprintf("test_tag_%s", testRandomSuffix)
//	tagCommit, err := GitTag(absPath, tagName, "test_tag_message", &sig)
//	require.NoError(t, err)
//	require.Equal(t, "0e06c18d12c415e9e2fd8eaf839b48d032cde1f1", tagCommit)
//
//	//test - push tag
//	//NOTE: if we use master as the remote branch name, this PR would be merged.
//	err = GitPush(absPath, localBranchName, gitRemote, fmt.Sprintf("remoteBranchName_%s", testRandomSuffix), tagName)
//
//	//assert
//	require.NoError(t, err)
//}
//
//func randomString(length int) string {
//	rand.Seed(time.Now().UnixNano())
//	b := make([]byte, length+2)
//	rand.Read(b)
//	return fmt.Sprintf("%x", b)[2 : length+2]
//}
