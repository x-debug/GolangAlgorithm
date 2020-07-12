package sort

import "awesomeProject/datastructs"

func InsertSort(arr []*datastructs.SortedElement) {
	if len(arr) < 2 {
		return
	}
	//从0~i就是有序节点
	for i := 0; i < len(arr)-1; i++ {
		k := i + 1 //无序的节点列表
		j := i

		//挨个比较有序列表
		for ; j >= 0 && arr[k].Key < arr[j].Key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = arr[k]
	}
}
