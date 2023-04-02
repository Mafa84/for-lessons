package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 4
	b := 6

	if err := For39(a, b); err != nil {
		return
	}
}

func For39(a int, b int) error {

	if a > b {
		return errors.New("а должен быть меньше b")
	}

	for i := a; i <= b; i++ {
		fmt.Println("i = ", i)
		for j := 0; j < i; j++ {
			fmt.Println(i)
		}
	}

	return nil
}
