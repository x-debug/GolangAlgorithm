package sort

import "awesomeProject/datastructs"

func InsertSort(arr []*datastructs.SortedElement) {
	if len(arr) < 2 {
		return
	}
	//从0~i就是有序节点
	for i := 1; i < len(arr); i++ {
		j := i - 1
		value := arr[i]
		//挨个比较有序列表
		for ; j >= 0 && value.Key < arr[j].Key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = value
	}
}
