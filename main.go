package main

import "fmt"

func main() {
	c, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", c)
}
