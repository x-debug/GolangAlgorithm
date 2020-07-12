package sort

import (
	"awesomeProject/datastructs"
	"fmt"
)

func MergeSort(arr []*datastructs.SortedElement) {
	fmt.Printf("pointer to %p", &arr)
	rMergeSort(arr, 0, len(arr)-1)
}

func rMergeSort(arr []*datastructs.SortedElement, p int, r int) {
	if p >= r {
		return
	}

	//求中间位置元素
	q := (p + r) / 2

	//排序p~q元素
	rMergeSort(arr, p, q)

	//排序q+1~r元素
	rMergeSort(arr, q+1, r)
	merge(arr, p, q, r)
}

func merge(arr []*datastructs.SortedElement, p int, q int, r int) {
	i := p
	j := q + 1
	k := 0
	temp := make([]*datastructs.SortedElement, r-p+1)

	for ; i <= q && j <= r; {
		if arr[i].Key < arr[j].Key {
			temp[k] = arr[i]
			k++
			i++
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}

	var begin int = i
	var end int = q

	if j <= r {
		begin = j
		end = r
	}

	for i = begin; i <= end; i++ {
		temp[k] = arr[i]
		k++
	}

	k = 0
	for i = p; i <= r; i++ {
		arr[i] = temp[k]
		k++
	}
}
