package datastructs

import (
	"fmt"
	"strings"
	"testing"
)

func doublyCompLinkedList(a interface{}, b interface{}) bool {
	sa, sb := a.(string), b.(string)
	return strings.Compare(sa, sb) == 0
}

func TestCreateDoublyLinkedList(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)

	if sll.size != 0 {
		t.Error("个数出现了问题")
	}
}

func TestDoublyLinkedList_AddFront(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddFront("hello")
	sll.AddFront("world")
	sll.AddFront("google")
	sll.AddFront("sorry")
	sll.AddFront("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
}

func TestDoublyLinkedList_AddEnd(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("baidu")

	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
}

func TestDoublyLinkedList_AddBefore(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("baidu")

	//hello world google sorry baidu
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)

	sll.AddBefore(sll.head, "fuck")
	sll.AddBefore(sll.head, "man")
	sll.AddBefore(sll.head, "women")

	if sll.Size() != 8 {
		t.Error("个数出现了问题")
	}

	fmt.Println("============================")
	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
}

func TestDoublyLinkedList_AddAfter(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("baidu")

	//hello world google sorry baidu
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)

	sll.AddAfter(sll.tail, "fuck")
	sll.AddAfter(sll.tail, "man")
	sll.AddAfter(sll.tail, "women")

	if sll.Size() != 8 {
		t.Error("个数出现了问题")
	}

	fmt.Println("============================")
	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
	doubleOrderNodes([]string{"hello", "world", "google", "sorry", "baidu", "fuck", "man", "women"}, sll, t)
}

func TestDoublyLinkedList_DeleteNode(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("baidu")

	//hello world google sorry baidu
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)

	sll.DeleteNode(sll.head)
	sll.DeleteNode(sll.tail)
	sll.DeleteNode(sll.tail)

	if sll.Size() != 2 {
		t.Error("个数出现了问题")
	}

	fmt.Println("============================")
	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
	doubleOrderNodes([]string{"world", "google", "sorry"}, sll, t)
}

func TestDoublyLinkedList_SearchForward(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("world")
	sll.AddEnd("baidu")

	//hello world google sorry baidu
	if sll.Size() != 6 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)

	node := sll.SearchBackward("world")
	sll.DeleteNode(node)

	//hello world google sorry baidu
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}
	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
	doubleOrderNodes([]string{"hello", "google", "sorry", "world", "baidu"}, sll, t)
}

func TestDoublyLinkedList_SearchBackward(t *testing.T) {
	sll := CreateDoublyLinkedList(doublyCompLinkedList)
	sll.AddEnd("hello")
	sll.AddEnd("world")
	sll.AddEnd("google")
	sll.AddEnd("sorry")
	sll.AddEnd("world")
	sll.AddEnd("baidu")

	//hello world google sorry baidu
	if sll.Size() != 6 {
		t.Error("个数出现了问题")
	}

	doublyLookup(sll, t)
	doublyLookNodes(sll, t)

	node := sll.SearchForward("world")
	sll.DeleteNode(node)

	//hello world google sorry baidu
	if sll.Size() != 5 {
		t.Error("个数出现了问题")
	}
	doublyLookup(sll, t)
	doublyLookNodes(sll, t)
	doubleOrderNodes([]string{"hello", "world", "google", "sorry", "baidu"}, sll, t)
}

//此函数为了测试链表的连续性
func doublyLookup(sll *DoublyLinkedList, t *testing.T) {
	var i int = 0

	for p := sll.head; p != nil; p = p.next {
		i++
	}

	if i != sll.Size() {
		t.Errorf("迭代不连续, i:%d, Size():%d", i, sll.Size())
	}
}

func doublyLookNodes(sll *DoublyLinkedList, t *testing.T) {
	for p := sll.head; p != nil; p = p.next {
		fmt.Println(p.data)
	}
}

func doubleOrderNodes(right []string, sll *DoublyLinkedList, t *testing.T) {
	var i int = 0
	for p := sll.head; p != nil; p = p.next {
		if right[i] != p.data {
			t.Error("元素次序不对")
		}

		i++
	}
}
