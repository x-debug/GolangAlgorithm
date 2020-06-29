package datastructs

import (
	"fmt"
	"strings"
	"testing"
)

func compLinkedList(a interface{}, b interface{}) bool {
	sa, sb := a.(string), b.(string)
	return strings.Compare(sa, sb) == 0
}

func TestCreateLinkedList(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)

	if sll.size != 0 {
		t.Error("个数出现了问题")
	}
}

func TestSinglyLinkedList_AddAtFront(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtFront("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtFront("sorry")
	sll.AddAtFront("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_AddAtFront2(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtFront("hello")
	sll.AddAtFront("world")
	sll.AddAtEnd("google")
	sll.AddAtEnd("sorry")
	sll.AddAtFront("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_AddAtFront3(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtFront("hello")

	if sll.Size() != 1 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_AddAtEnd(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtEnd("world")
	sll.AddAtEnd("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_AddAtEnd2(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtEnd("hello")
	sll.AddAtFront("world")
	sll.AddAtFront("google")
	sll.AddAtEnd("sorry")
	sll.AddAtEnd("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_AddAtEnd3(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
	sll.AddAtEnd("hello")

	if sll.Size() != 1 {
		t.Error("个数出现了问题")
	}

	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_DeleteNode(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
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
	lookup(sll, t)
	lookNodes(sll, t)
}

func TestSinglyLinkedList_Search(t *testing.T) {
	sll := CreateLinkedList(compLinkedList)
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

	lookup(sll, t)
	lookNodes(sll, t)
}

//此函数为了测试链表的连续性
func lookup(sll *SinglyLinkedList, t *testing.T) {
	var i int = 0

	for p := sll.head; p != nil; p = p.next {
		i++
	}

	if i != sll.Size() {
		t.Errorf("迭代不连续, i:%d, Size():%d", i, sll.Size())
	}
}

func lookNodes(sll *SinglyLinkedList, t *testing.T) {
	for p := sll.head; p != nil; p = p.next {
		fmt.Println(p.data)
	}
}
