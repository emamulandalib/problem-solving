package main

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) (result []int) {
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}

	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return
}
