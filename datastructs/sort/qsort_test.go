package sort

import (
	"awesomeProject/datastructs"
	"testing"
)

func TestQSort(t *testing.T) {
	arr := []*datastructs.SortedElement{
		{Key: 3},
		{Key: 2},
		{Key: 6},
		{Key: 5},
		{Key: 4},
		{Key: 11},
		{Key: 1},
		{Key: 12},
		{Key: 31},
		{Key: 21},
		{Key: 17},
		{Key: 22},
		{Key: 66},
		{Key: 7},
		{Key: 90},
		{Key: 32},
		{Key: 22},
		{Key: 4},
	}

	QSort(arr)
	if !checkSorted(arr) {
		t.Error("TestQSort")
	}
}
