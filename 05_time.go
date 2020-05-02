package main

import (
	"fmt"
	"time"
)

func main() {

	var counter int8 = 1
	// while true loop
	for {
		fmt.Println("Повтор номер: ", counter)
		counter += 1
		time.Sleep(time.Second * 5)
	}

	// Alternative Version
	/*
		for true {
			fmt.Println("Infinite Loop 2")
			time.Sleep(time.Second)
		}
	*/
}
