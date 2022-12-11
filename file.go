package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var (
	ErrEmptyMemo = errors.New("memo name cannot be empty")
	ErrMakeMemo  = errors.New("failed to create memo")
)

func createMemo() (f string, err error) {
	fmt.Print("Please enter the name of the memo you want to create: ")
	s.Scan()
	l := s.Text()
	if l == "" {
		return "", ErrEmptyMemo
	}

	if !exists("LIFE") {
		err := os.Mkdir("LIFE", 0755)
		if err != nil {
			return "", ErrMakeMemo
		}
	}

	f = "LIFE/" + l + ".md"
	fl, err := os.Create(f)
	if err != nil {
		return "", ErrMakeMemo
	}
	defer fl.Close()

	return f, nil
}

func life(e string, f string) (err error) {
	_, err = exec.Command(e, f).Output()
	if err != nil {
		return err
	}
	return nil
}
