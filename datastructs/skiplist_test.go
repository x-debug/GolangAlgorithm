package datastructs

import (
	"fmt"
	"testing"
)

func TestCreateSkipList(t *testing.T) {
	sl := CreateSkipList()

	sl.Insert(5, "hello")
	sl.Insert(1, "world")
	sl.Insert(3, "test")
	sl.Insert(15, "some")
	sl.Insert(52, "fuck")
	sl.Insert(8, "hello1")

	b, value := sl.Search(3)
	if !b || value != "test" {
		t.Error("TestCreateSkipList")
	}
}

func TestSkipList_Insert(t *testing.T) {
	sl := CreateSkipList()

	sl.Insert(5, "hello")
	sl.Insert(1, "world")
	sl.Insert(3, "test")
	sl.Insert(15, "some")
	sl.Insert(52, "fuck")
	sl.Insert(8, "hello1")

	if sl.Size != 6 {
		t.Error("TestSkipList_Insert")
	}
}

func TestSkipList_Remove(t *testing.T) {
	sl := CreateSkipList()

	sl.Insert(5, "hello")
	sl.Insert(1, "world")
	sl.Insert(3, "test")
	sl.Insert(15, "some")
	sl.Insert(52, "fuck")
	sl.Insert(8, "hello1")

	sl.Remove(3)
	b, value := sl.Search(3)
	if b {
		t.Error(value)
	}
}

func TestSkipList_Search(t *testing.T) {
	sl := CreateSkipList()

	sl.Insert(5, "hello")
	sl.Insert(1, "world")

	_, value := sl.Search(2)
	fmt.Println(value)
}
