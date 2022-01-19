package main

func BinarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if target == nums[0] {
			return 1
		}
		return -1
	}

	return binarySearch(nums, target)
}

// arr = 3, 4, 6, 7 t = 7
func binarySearch(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		mid := (l + h) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			h = mid - 1
		}

		if target > nums[mid] {
			l = mid + 1
		}
	}

	return -1
}
