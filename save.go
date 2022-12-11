package main

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrInitGit = errors.New("git init failed")
	ErrAddGit  = errors.New("git add failed")
	ErrCmtGit  = errors.New("git commit failed")
	ErrPushGit = errors.New("git push failed")
)

func Save(g string) (err error) {
	fmt.Print("Please enter the name of the file you want to save: ")
	s.Scan()
	l := s.Text()
	if l == "" {
		return ErrEmptyMemo
	}
	f := "LIFE/" + l
	if err := gitinit(); err != nil {
		return err
	}
	_, err = exec.Command("git", "add", f).Output()
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
