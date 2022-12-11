package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("args is 2 or more required")
		os.Exit(1)
	}

	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "edit":
		if err := EditConfig(config.Editor); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "show":
		if err := Show(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "make":
		f, err := createMemo()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := life(config.Editor, f); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "save":
		if err := Save(config.GitURI); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("invalid command")
	}

}
