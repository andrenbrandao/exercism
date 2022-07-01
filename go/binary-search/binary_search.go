package binarysearch

func SearchInts(list []int, key int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (right-left)/2 + left

		if key == list[mid] {
			return mid
		} else if key < list[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
