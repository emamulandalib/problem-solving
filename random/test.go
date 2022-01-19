package main

import "fmt"

func RandTest() {
	fmt.Println("Please input a number:")
	var n int
	_, _ = fmt.Scanln(&n)

	step := 2
	der := 1

	for n <= 35 {
		fmt.Println(n)

		if der == 1 {
			step *= 2
		} else {
			step /= 2
		}

		if step > 16 || step < 2 {
			step = 4
			der *= -1
		}

		n += step * der
	}
}
