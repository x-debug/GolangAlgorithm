package leetcode

import (
	"awesomeProject/datastructs"
	"testing"
)

func TestP234MethodArray(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(4)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(3)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if !P234MethodArray(sll.Size(), sll.Head) {
		t.Error("P234MethodArray")
	}
}

func TestP234MethodArray2(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(7)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(6)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if P234MethodArray(sll.Size(), sll.Head) {
		t.Error("TestP234MethodArray2")
	}
}

func TestP234MethodRecursion(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(4)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(3)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if !P234MethodRecursion(sll.Size(), sll.Head) {
		t.Error("TestP234MethodRecursion")
	}
}

func TestP234MethodRecursion2(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(7)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(6)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if P234MethodRecursion(sll.Size(), sll.Head) {
		t.Error("TestP234MethodRecursion2")
	}
}

func TestP234MethodTwoPointer(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(4)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(3)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if !P234MethodTwoPointer(sll.Size(), sll.Head) {
		t.Error("TestP234MethodTwoPointer")
	}
}

func TestP234MethodTwoPointer2(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(3)
	sll.AddAtEnd(7)
	sll.AddAtEnd(5)
	sll.AddAtEnd(4)
	sll.AddAtEnd(6)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if P234MethodTwoPointer(sll.Size(), sll.Head) {
		t.Error("TestP234MethodTwoPointer2")
	}
}

func TestP234MethodTwoPointer3(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(0)
	sll.AddAtEnd(1)

	if !P234MethodTwoPointer(sll.Size(), sll.Head) {
		t.Error("TestP234MethodTwoPointer2")
	}
}

func TestP234MethodTwoPointer4(t *testing.T) {
	sll := datastructs.CreateLinkedList(nil)
	sll.AddAtEnd(1)
	sll.AddAtEnd(2)
	sll.AddAtEnd(2)
	sll.AddAtEnd(1)

	if !P234MethodTwoPointer(sll.Size(), sll.Head) {
		t.Error("TestP234MethodTwoPointer2")
	}
}
