package main

func pickFromBotSide(A []int, B int) int {
	sum, i := 0, 0

	for i < B {
		sum += A[i]
	}

	return sum
}
