package main

import (
	"bufio"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
)

const yaml_name = "life.yaml"

var s = bufio.NewScanner(os.Stdin)

var (
	ErrCreateConfig = errors.New("failed to create config file")
	ErrEmptyEditor  = errors.New("editor cannot be empty")
	ErrEmptyGitUri  = errors.New("git uri cannot be empty")
	ErrOpenConfig   = errors.New("failed to open config file")
	ErrWriteConfig  = errors.New("failed to write to config file")
	ErrReadConfig   = errors.New("failed to read config file")
	ErrUnmarshal    = errors.New("failed to unmarshal config file")
	ErrEditConfig   = errors.New("failed to edit config file")
)

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
		return ErrCreateConfig
	}
	defer f.Close()
	return nil
}

func writeEditor() (string, error) {
	fmt.Print("Please enter your editor of choice(ex: vi,vim,nvim,code):\n> ")
	s.Scan()
	editor := s.Text()
	if editor == "" {
		return "", ErrEmptyEditor
	}
	f, err := os.OpenFile(yaml_name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return "", ErrOpenConfig
	}
	defer f.Close()
	if _, err = f.WriteString("editor: " + editor + "\n"); err != nil {
		return "", ErrWriteConfig
	}
	return editor, nil
}

func writeGitUri() (string, error) {
	fmt.Print("Please enter your git uri:\n> ")
	s.Scan()
	git_uri := s.Text()
	if git_uri == "" {
		return "", ErrEmptyGitUri
	}
	f, err := os.OpenFile(yaml_name, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return "", ErrOpenConfig
	}
	defer f.Close()
	if _, err = f.WriteString("git_uri: " + git_uri + "\n"); err != nil {
		return "", ErrWriteConfig
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
		return res, ErrReadConfig
	}
	d := make(map[string]string)
	if err := yaml.Unmarshal(y, &d); err != nil {
		return res, ErrUnmarshal
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

func EditConfig(e string) (err error) {
	if !exists(yaml_name) {
		if err := createYaml(); err != nil {
			return err
		}
	}
	_, err = exec.Command(e, yaml_name).Output()
	if err != nil {
		return ErrEditConfig
	}
	return nil
}
