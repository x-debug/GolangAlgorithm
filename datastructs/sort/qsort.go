package sort

import "awesomeProject/datastructs"

func QSort(arr []*datastructs.SortedElement) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []*datastructs.SortedElement, b int, e int) {
	if b >= e {
		return
	}

	pivot := arr[e]
	var i int = b // i 左边都是有序(包括i),i右边都是无序

	for j := b; j < e; j++ {
		if arr[j].Key < pivot.Key {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[e] = arr[e], arr[i]
	quickSort(arr, b, i-1)
	quickSort(arr, i+1, e)
}
