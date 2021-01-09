package repo

import (
	"github.com/go-git/go-git/v5"
	"os"
)

func ShallowClone(path, gitRepo string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		Depth: 1,
		URL:      gitRepo,
		Progress: os.Stdout,
	})
	if err != nil {
		return  err
	}
	return nil
}

