package git

import git2go "gopkg.in/libgit2/git2go.v25"

func GitGetBranch(repoPath string) (string, error) {
	repo, oerr := git2go.OpenRepository(repoPath)
	if oerr != nil {
		return "", oerr
	}

	currentBranch, berr := repo.Head()
	if berr != nil {
		return "", berr
	}

	return currentBranch.Branch().Name()
}
