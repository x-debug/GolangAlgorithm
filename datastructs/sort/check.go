package sort

import "awesomeProject/datastructs"

func checkSorted(arr []*datastructs.SortedElement) bool {
	if len(arr) < 2 {
		return true
	}

	var prev = arr[0]

	for i := 1; i < len(arr); i++ {
		if prev.Key > arr[i].Key {
			return false
		}
		prev = arr[i]
	}

	return true
}
