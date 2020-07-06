package datastructs

import (
	"fmt"
	"strings"
	"testing"
)

func compCircularLinkedList(a interface{}, b interface{}) bool {
	sa, sb := a.(string), b.(string)
	return strings.Compare(sa, sb) == 0
}

func TestCreateCircularLinkedList(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)

	if sll.size != 0 {
		t.Error("个数出现了问题")
	}
}

func TestCircularLinkedList_AddAtFront(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtFront("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtFront("sorry")
	sll.AddAtFront("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_AddAtFront2(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtFront("hello")
	sll.AddAtFront("world")
	sll.AddAtEnd("google")
	sll.AddAtEnd("sorry")
	sll.AddAtFront("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_AddAtFront3(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtFront("hello")

	if sll.Size() != 1 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_AddAtEnd(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtEnd("world")
	sll.AddAtEnd("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_AddAtEnd2(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_AddAtEnd3(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")

	if sll.Size() != 1 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_DeleteNode(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	sll.DeleteNode("sorry")
	sll.DeleteNode("sorry123")
	if sll.Size() != 4 {
		t.Error("个数出现了问题")
	}

	sll.AddAtFront("fuck")
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_DeleteNode2(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	//google world hello sorry baidu

	sll.DeleteNode("google")
	sll.DeleteNode("sorry123")
	if sll.Size() != 4 {
		t.Error("个数出现了问题")
	}

	sll.AddAtFront("fuck")
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	sll.DeleteNode("fuck")
	if sll.Size() != 4 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

func TestCircularLinkedList_Search(t *testing.T) {
	sll := CreateCircularLinkedList(compCircularLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	data := sll.Search("sorry")

	if data.(string) != "sorry" {
		t.Error("数据出现了问题")
	}

	sll.AddAtFront("fuck")
	if sll.Size() != 6 {
		t.Error("个数出现了问题")
	}

	lookupCircular(sll, t)
	lookNodesCircular(sll, t)
}

//此函数为了测试链表的连续性
func lookupCircular(sll *CircularLinkedList, t *testing.T) {
	var i int = 0

	for p := sll.Head; p != sll.Tail; p = p.Next {
		i++
	}

	if i+1 != sll.Size() {
		t.Errorf("迭代不连续, i:%d, Size():%d", i+1, sll.Size())
	}
}

func lookNodesCircular(sll *CircularLinkedList, t *testing.T) {
	index := 0
	for p := sll.Head; index < sll.size; p = p.Next {
		fmt.Println(p.Data)
		index++
	}
}
