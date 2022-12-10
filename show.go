package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func show() error {
	fs, err := getFiles("LIFE")
	if err != nil {
		return err
	}
	for _, f := range fs {
		fmt.Println(f)
	}
	return nil
}

func getFiles(path string) ([]string, error) {
	fs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, f := range fs {
		if f.IsDir() {
			d, err := getFiles(filepath.Join(path, f.Name()))
			if err != nil {
				return nil, err
			}
			files = append(files, d...)
			continue
		}
		files = append(files, filepath.Join(path, f.Name()))
	}
	return files, nil
}
