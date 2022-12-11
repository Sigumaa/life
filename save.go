package main

import (
	"errors"
	"os/exec"
)

func Save(g string, l string) (err error) {
	if err := gitinit(); err != nil {
		return err
	}
	_, err = exec.Command("git", "add", l).Output()
	if err != nil {
		return errors.New("git add failed")
	}
	cm := "life: " + l
	_, err = exec.Command("git", "commit", "-m", cm).Output()
	if err != nil {
		return errors.New("git commit failed")
	}
	_, err = exec.Command("git", "push", g).Output()
	if err != nil {
		return errors.New("git push failed")
	}
	return nil
}

func gitinit() (err error) {
	_, err = exec.Command("git", "status").Output()
	if err != nil {
		_, err = exec.Command("git", "init").Output()
		if err != nil {
			return errors.New("git init failed")
		}
	}
	return nil
}
