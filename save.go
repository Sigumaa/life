package main

import (
	"errors"
	"os/exec"
)

var (
	ErrInitGit = errors.New("git init failed")
	ErrAddGit  = errors.New("git add failed")
	ErrCmtGit  = errors.New("git commit failed")
	ErrPushGit = errors.New("git push failed")
)

func Save(g string, l string) (err error) {
	if err := gitinit(); err != nil {
		return err
	}
	_, err = exec.Command("git", "add", l).Output()
	if err != nil {
		return ErrAddGit
	}
	cm := "life: " + l
	_, err = exec.Command("git", "commit", "-m", cm).Output()
	if err != nil {
		return ErrCmtGit
	}
	_, err = exec.Command("git", "push", g).Output()
	if err != nil {
		return ErrPushGit
	}
	return nil
}

func gitinit() (err error) {
	_, err = exec.Command("git", "status").Output()
	if err != nil {
		_, err = exec.Command("git", "init").Output()
		if err != nil {
			return ErrInitGit
		}
	}
	return nil
}
