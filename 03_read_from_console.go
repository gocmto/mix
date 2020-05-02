package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("Как тебя зовут?")
	fmt.Scanf("%s\n", &name)

	/*
		if len(mystring) > 0 { }
		if mystring != "" { }
	*/

	if len(name) == 0 {
		fmt.Println("Нечестно! Как тебя зовут?")
		fmt.Scanf("%s\n", &name)
	} else {
		var age int
		fmt.Println("Сколько тебе лет?")
		fmt.Scanf("%d\n", &age)

		fmt.Printf("Привет, %s, твой возраст - %d\n", name, age)
	}

	var quit string
	fmt.Scanf("%v", &quit)

}
