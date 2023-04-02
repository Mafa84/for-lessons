package exercises

import (
	"errors"
	"fmt"
)

/*
	 Даны целые положительные числа A и B (A < B). Вывести все целые
	числа от A до B включительно; при этом каждое число должно выводиться
	столько раз, каково его значение (например, число 3 выводится 3 раза).
*/

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
