package git

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

func GitSignature(authorName string, authorEmail string) *object.Signature {
	return &object.Signature{
		Name:  authorName,
		Email: authorEmail,
		When:  time.Now(),
	}
}
