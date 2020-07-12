package sort

import (
	"awesomeProject/datastructs"
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []*datastructs.SortedElement{
		{Key: 3},
		{Key: 2},
		{Key: 6},
		{Key: 5},
		{Key: 4},
		{Key: 1},
	}

	fmt.Printf("pointer to %p\r\n", &arr)
	MergeSort(arr)
	if !checkSorted(arr) {
		t.Error("TestMergeSort")
	}
}
