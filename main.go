package main

import "fmt"

func main() {
	res, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	fmt.Println("Editing config...")
	if err := EditConfig(); err != nil {
		panic(err)
	}
}
