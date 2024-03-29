package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/packagrio/go-common/pipeline"
)

func GitTag(repoPath string, version string, message string, signature *object.Signature) (string, error) {
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
	commit := commitObject.Hash

	//tagId, terr := repo.Tags.CreateLightweight(version, commit, false)
	tagId, terr := repo.CreateTag(version, commit, &git.CreateTagOptions{
		Tagger:  signature,
		Message: fmt.Sprintf("(%s) %s", version, message),
	})
	if terr != nil {
		return "", terr
	}

	tagObj, terr := repo.TagObject(tagId.Hash())
	if terr != nil {
		return "", terr
	}
	return tagObj.Target.String(), terr
}

func GitGetTagDetails(repoPath string, tagName string) (*pipeline.GitTagDetails, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return nil, oerr
	}

	tagRef, terr := repo.Tag(tagName)
	if terr != nil {
		return nil, terr
	}

	tag, lerr := repo.TagObject(tagRef.Hash()) //assume its an annotated tag.

	var currentTag *pipeline.GitTagDetails
	if lerr != nil {
		//this is a lightweight tag, not an annotated tag.
		commitRef, rerr := repo.CommitObject(tagRef.Hash())
		if rerr != nil {
			return nil, rerr
		}

		author := commitRef.Author

		log.Printf("Light-weight tag (%s) Commit ID: %s, DATE: %s", tagName, commitRef.Hash.String(), author.When.String())

		currentTag = &pipeline.GitTagDetails{
			TagShortName: tagName,
			CommitSha:    commitRef.Hash.String(),
			CommitDate:   author.When,
		}

	} else {

		log.Printf("Annotated tag (%s) Tag ID: %s, Commit ID: %s, DATE: %s", tagName, tag.Hash.String(), tag.Target.String(), tag.Tagger.When.String())

		currentTag = &pipeline.GitTagDetails{
			TagShortName: tagName,
			CommitSha:    tag.Target.String(),
			CommitDate:   tag.Tagger.When,
		}
	}
	return currentTag, nil
}

// Get the nearest tag on branch.
// tag must be nearest, ie. sorted by their distance from the HEAD of the branch, not the date or tagname.
// basically `git describe --tags --abbrev=0`
// https://github.com/go-git/go-git/pull/584
func GitFindNearestTagName(repoPath string) (string, error) {
	repo, oerr := git.PlainOpen(repoPath)
	if oerr != nil {
		return "", oerr
	}

	// We assume HEAD is the commit of the new release and already tagged by packagr-releasr,
	// thus the search for nearest tag is started from HEAD^
	// see also https://github.com/PackagrIO/go-common/commit/867857047ac8f592b6fe64fde5630f38876b2656
	parentCommit, err := repo.ResolveRevision(plumbing.Revision("HEAD^"))
	if err != nil {
		return "", fmt.Errorf("could not get HEAD^ commit: %v", err)
	}
	logIter, err := repo.Log(&git.LogOptions{
		From:  *parentCommit,
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return "", fmt.Errorf("could not get log: %v", err)
	}

	tags, err := buildTagRefMap(repo)
	if err != nil {
		return "", fmt.Errorf("could not build tag ref map: %v", err)
	}

	var tagStr string
	logIter.ForEach(func(c *object.Commit) error {
		if tag, exists := tags[c.Hash]; exists {
			tagStr = tag.Name().Short()
			return storer.ErrStop
		}
		return nil
	})

	if tagStr != "" {
		return tagStr, nil
	}
	return "", fmt.Errorf("could not find latest tag")
}

// from https://github.com/go-git/go-git/pull/584/files
func buildTagRefMap(r *git.Repository) (map[plumbing.Hash]*plumbing.Reference, error) {
	iter, err := r.Tags()
	if err != nil {
		return nil, err
	}
	tags := map[plumbing.Hash]*plumbing.Reference{}

	if err := iter.ForEach(func(ref *plumbing.Reference) error {
		// log.Printf("Tag: %s", ref.Name())
		obj, err := r.TagObject(ref.Hash())
		switch err {
		case nil:
			// Tag object present
			// t is an annotated tag
			tags[obj.Target] = ref
		case plumbing.ErrObjectNotFound:
			// t is a lighweight tag
			tags[ref.Hash()] = ref
		default:
			// Some other error
			return err
		}
		return nil
	}); err != nil {
		// Handle outer iterator error
		return nil, err
	}

	return tags, nil
}
