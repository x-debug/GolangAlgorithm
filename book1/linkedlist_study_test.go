package book1

import (
	"awesomeProject/datastructs"
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	ll := datastructs.CreateLinkedList(nil)

	ll.AddAtEnd("A")
	ll.AddAtEnd("B")
	ll.AddAtEnd("C")
	ll.AddAtEnd("D")
	ll.AddAtEnd("E")
	ll.AddAtEnd("F")

	testOrderNodes([]string{"A", "B", "C", "D", "E", "F"}, ll, t)
	ReverseLinkedList(ll)
	lookNodes(ll, t)
	testOrderNodes([]string{"F", "E", "D", "C", "B", "A"}, ll, t)
}

func TestJoinTwoLinkedList1(t *testing.T) {
	ll1 := datastructs.CreateLinkedList(nil)

	ll1.AddAtEnd(1)
	ll1.AddAtEnd(2)
	ll1.AddAtEnd(3)
	ll1.AddAtEnd(4)
	ll1.AddAtEnd(5)

	ll2 := datastructs.CreateLinkedList(nil)
	ll2.AddAtEnd(2)
	ll2.AddAtEnd(5)
	ll2.AddAtEnd(7)
	ll2.AddAtEnd(8)
	ll2.AddAtEnd(9)
	ll2.AddAtEnd(15)

	r := JoinTwoLinkedList1(ll1, ll2)
	lookNodes(r, t)
	testOrderNodesInt([]int{1, 2, 2, 3, 4, 5, 5, 7, 8, 9, 15}, r, t)
}

func TestGetLinkedMiddleNode(t *testing.T) {
	ll1 := datastructs.CreateLinkedList(nil)

	ll1.AddAtEnd(1)
	ll1.AddAtEnd(2)
	ll1.AddAtEnd(3)
	ll1.AddAtEnd(4)
	ll1.AddAtEnd(5)

	p := GetLinkedMiddleNode(ll1)
	fmt.Println(p.Data)
}

func TestJoinTwoLinkedList2(t *testing.T) {
	ll1 := datastructs.CreateLinkedList(nil)

	ll1.AddAtEnd(1)
	ll1.AddAtEnd(2)
	ll1.AddAtEnd(3)
	ll1.AddAtEnd(4)
	ll1.AddAtEnd(5)
	ll1.AddAtEnd(6)

	p := GetLinkedMiddleNode(ll1)
	fmt.Println(p.Data)
}

func TestDeleteLinkedNode(t *testing.T) {
	ll1 := datastructs.CreateLinkedList(nil)

	ll1.AddAtEnd(1)
	ll1.AddAtEnd(2)
	ll1.AddAtEnd(3)
	ll1.AddAtEnd(4)
	ll1.AddAtEnd(5)
	ll1.AddAtEnd(6)
	ll1.AddAtEnd(7)
	ll1.AddAtEnd(8)

	DeleteLinkedNode(ll1, 8)
	lookNodes(ll1, t)
}

func TestCheckCircleLinked(t *testing.T) {
	ll1 := datastructs.CreateLinkedList(nil)

	ll1.AddAtEnd(1)
	ll1.AddAtEnd(2)
	//ll1.AddAtEnd(3)
	//ll1.Tail.Next = ll1.Head

	fmt.Println(CheckCircleLinked(ll1))
}

func testOrderNodes(right []string, sll *datastructs.SinglyLinkedList, t *testing.T) {
	var i int = 0
	for p := sll.Head; p != nil; p = p.Next {
		if right[i] != p.Data {
			t.Error("元素次序不对")
		}

		i++
	}
}

func testOrderNodesInt(right []int, sll *datastructs.SinglyLinkedList, t *testing.T) {
	var i int = 0
	for p := sll.Head; p != nil; p = p.Next {
		if right[i] != p.Data {
			t.Error("元素次序不对")
		}

		i++
	}
}

func lookNodes(sll *datastructs.SinglyLinkedList, t *testing.T) {
	for p := sll.Head; p != nil; p = p.Next {
		fmt.Println(p.Data)
	}
}
