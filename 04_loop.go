package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Press 1 to run")
	fmt.Println("Press 2 to exit")
	for {
		sample()
	}
}

func sample() {
	var input int
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")
		return
	}
	switch input {
	case 1:
		fmt.Println("hi")
	case 2:
		os.Exit(2)
	default:
		fmt.Println("def")
	}
}
