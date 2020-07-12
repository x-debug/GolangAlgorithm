package sort

import "awesomeProject/datastructs"

func SelectSort(arr []*datastructs.SortedElement) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j].Key < arr[min].Key {
				min = j
			}
		}

		arr[min], arr[i] = arr[i], arr[min]
	}
}
