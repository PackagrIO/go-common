package git

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGitGenerateChangelog(t *testing.T) {

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
	changelog, err := GitGenerateChangelog(absPath,
		"4feed9fb27bfebba92d18839fb0a19866b7eb16a", //(v1.0.10) Automated packaging of release by CapsuleCD
		"c800faadaed8ad71f3ddf1fb4bc3f22c6d8969a1", //Merge 2ec74dd into 4feed9f
	)

	//assert
	require.NoError(t, err)
	require.Equal(t, `Timestamp |  SHA | Message | Author
------------- | ------------- | ------------- | -------------
2019-10-03T15:46Z | c800faad | Merge 2ec74dd9542e5be652195823aaf27d9579c60497 into 4feed9fb27bfebba92d18839fb0a19866b7eb16a | Jason Kulatunga
2019-10-03T15:45Z | 2ec74dd9 | added index.js file for testing. copied from left-pad. | Jason Kulatunga
2019-10-03T15:41Z | 020f4eef | added eslint | Jason Kulatunga
2019-10-03T15:40Z | 21ec5e82 | adding eslint config | Jason Kulatunga
2019-10-03T15:35Z | f6f820c1 | Update README.md | Jason Kulatunga
`, changelog)
}

// Validate handling of commit messages from PR squash merges.
// Github creates squash commits with cr/lf line endings and creates a list of all squashed commit messages using asterisk list headers.
func TestCleanCommitMessageFromSquash(t *testing.T) {

	//setup
	squashCommit := "Commit a (#13)\n\n* Commit a\r\n\r\n* commit b\r\n\r\n* another commit"

	//test
	cleanedMsg := cleanCommitMessage(squashCommit)

	//assert
	require.Equal(t, "Commit a (#13)<li>Commit a<li>commit b<li>another commit", cleanedMsg)
}
