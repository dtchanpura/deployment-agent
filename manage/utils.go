package manage

import (
	"gopkg.in/src-d/go-git.v4"
)

func PullRepository(repoPath string, remoteName string) error {
	if remoteName == "" {
		remoteName = "origin"
	}
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: remoteName})
	if err != nil {
		return err
	}
	return nil
}
