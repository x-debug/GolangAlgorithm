package sort

import "awesomeProject/datastructs"

/**
冒泡排序思路:
比较两个相邻元素，前面元素比后面元素大，就交换两个元素
*/
func BubbleSort(arr []*datastructs.SortedElement) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i].Key > arr[j].Key {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
