package main

import (
	"fmt"
	"sort"
)

func uniqueArr(arr []int) []int {
	var uniqueArr []int
	var data map[int]bool

	for _, a := range arr {
		if _, ok := data[a]; !ok {
			data[a] = true
			uniqueArr = append(uniqueArr, a)
		}
	}

	return uniqueArr
}

func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}

	uniqueCoins := uniqueArr(coins)

	if len(uniqueCoins) == 1 {
		return uniqueCoins[0] + 1
	}

	sort.Ints(uniqueCoins)

	if uniqueCoins[0] == 1 {
		return 1
	}

	lastCoin := uniqueCoins[len(coins)-1]

	for i := 2; i <= lastCoin; i++ {

	}

	return 0
}

func main() {
	// 1 2 3 5 7 22
	fmt.Println(NonConstructibleChange([]int{5, 7, 1, 1, 2, 3, 22}))
}
