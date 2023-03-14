package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

//Add all modified files to index, and commit.
func GitCommit(repoPath string, message string, signature *object.Signature) error {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return oerr
	}

	//get repo working tree.
	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}
	aerr := workTree.AddWithOptions(&git.AddOptions{All: true})
	if aerr != nil {
		return aerr
	}

	// We can verify the current status of the worktree using the method Status.
	status, serr := workTree.Status()
	log.Printf("Status: %v", status)
	if serr != nil {
		return serr
	}

	// Commits the current staging area to the repository
	_, cerr := workTree.Commit(message, &git.CommitOptions{
		Author: signature,
	})

	return cerr
}

func GitGetHeadCommit(repoPath string) (string, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}
	commitHead, herr := repo.Head()
	if herr != nil {
		return "", herr
	}
	commitObject, lerr := repo.CommitObject(commitHead.Hash())
	if lerr != nil {
		return "", lerr
	}
	return commitObject.Hash.String(), nil
}
