package git

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGitSignature(t *testing.T) {
	//setup
	authorName := "test"
	authorEmail := "test@example.com"

	//test
	signature := GitSignature(authorName, authorEmail)

	//assert
	require.Equal(t, signature.Name, authorName)
	require.Equal(t, signature.Email, authorEmail)
}
