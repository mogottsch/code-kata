package search

func binarySearch(needle int, arr []int) int {
	start := 0
	end := len(arr)
	index := (end + start) / 2

	if len(arr) == 0 {
		return -1
	}

	for {
		if index >= len(arr) {
			return -1
		}

		element := arr[index]
		if element == needle {
			return index
		}

		if start == end {
			return -1
		}

		if element < needle {
			start = index + 1
		} else {
			end = index
		}
		index = (end + start) / 2
	}
}
