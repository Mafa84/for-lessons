package main

import (
	"fmt"
	exercises "forExer/funcs"
)

func main() {

	var a, b int
	var err error

	fmt.Println("Введите число А:")
	_, err = fmt.Scan(&a)
	if err != nil {
		fmt.Println("Scan error: ", err)
		return
	}

	fmt.Println("Введите число B:")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Scan error: ", err)
		return
	}

	err = exercises.For39(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
