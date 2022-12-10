package main

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const yaml_name = "life.yaml"

var s = bufio.NewScanner(os.Stdin)

type Config struct {
	Editor string
	GitURI string
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func createYaml() error {
	f, err := os.Create(yaml_name)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func writeEditor() (string, error) {
	fmt.Println("Please enter your editor of choice(vi,vim,nvim,vscode):\n> ")
	s.Scan()
	editor := s.Text()
	if editor == "" {
		return "", errors.New("editor cannot be empty")
	}
	f, err := os.OpenFile(yaml_name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err = f.WriteString("editor: " + editor); err != nil {
		return "", err
	}
	return editor, nil
}

func writeGitUri() (string, error) {
	fmt.Println("Please enter your git uri:\n> ")
	s.Scan()
	git_uri := s.Text()
	if git_uri == "" {
		return "", errors.New("git uri cannot be empty")
	}
	f, err := os.OpenFile(yaml_name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err = f.WriteString("git_uri: " + git_uri); err != nil {
		return "", err
	}
	return git_uri, nil
}

func LoadConfig() (res Config, err error) {
	if !exists(yaml_name) {
		if err := createYaml(); err != nil {
			return res, err
		}
	}
	y, err := os.ReadFile(yaml_name)
	if err != nil {
		return res, err
	}
	d := make(map[string]string)
	if err := yaml.Unmarshal(y, &d); err != nil {
		return res, err
	}
	e, ok := d["editor"]
	if !ok {
		e, err = writeEditor()
		if err != nil {
			return res, err
		}
	}
	g, ok := d["git-uri"]
	if !ok {
		g, err = writeGitUri()
		if err != nil {
			return res, err
		}
	}
	res.Editor = e
	res.GitURI = g
	return res, nil
}
