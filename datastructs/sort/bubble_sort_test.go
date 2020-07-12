package sort

import (
	"awesomeProject/datastructs"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []*datastructs.SortedElement{
		{Key: 3},
		{Key: 2},
		{Key: 6},
		{Key: 5},
		{Key: 4},
		{Key: 1},
	}

	BubbleSort(arr)
	if !checkSorted(arr) {
		t.Error("TestBubbleSort")
	}
}
