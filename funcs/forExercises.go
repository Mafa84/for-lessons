package exercises

import "fmt"

/*
	 Даны целые положительные числа A и B (A < B). Вывести все целые
	числа от A до B включительно; при этом каждое число должно выводиться
	столько раз, каково его значение (например, число 3 выводится 3 раза).
*/

func For39(a int, b int) error {

	for i := 0; i < a; i++ {
		fmt.Println(a)
	}

	return nil
}