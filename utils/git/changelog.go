package git

import (
	"fmt"
	goUtils "github.com/analogj/go-util/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"strings"
)

//https://github.com/go-git/go-git/issues/36
func GitGenerateChangelog(repoPath string, baseSha string, headSha string) (string, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}

	markdown := goUtils.StripIndent(`Timestamp |  SHA | Message | Author
	------------- | ------------- | ------------- | -------------
	`)

	start, err := repo.ResolveRevision(plumbing.Revision(baseSha))
	if err != nil {
		return "", err
	}

	end, err := repo.ResolveRevision(plumbing.Revision(headSha))
	if err != nil {
		return "", err
	}

	logIter, err := repo.Log(&git.LogOptions{
		From:  *end,
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return "", err
	}

	logIter.ForEach(func(c *object.Commit) error {
		if c.Hash == *start {
			return storer.ErrStop
		}

		markdown += fmt.Sprintf("%s | %.8s | %s | %s\n", //TODO: this should have a link for the SHA.
			c.Author.When.UTC().Format("2006-01-02T15:04Z"),
			c.Hash.String(),
			cleanCommitMessage(c.Message),
			c.Author.Name,
		)

		return nil
	})

	return markdown, nil
}

// helpers
func cleanCommitMessage(commitMessage string) string {
	commitMessage = strings.TrimSpace(commitMessage)
	if commitMessage == "" {
		return "--"
	}

	commitMessage = strings.Replace(commitMessage, "|", "/", -1)
	commitMessage = strings.Replace(commitMessage, "\n", " ", -1)

	return commitMessage
}
