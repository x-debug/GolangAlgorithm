package datastructs

type SinglyLinkedListNode struct {
	Next *SinglyLinkedListNode

	Data interface{}
}

type SinglyLinkedList struct {
	Head *SinglyLinkedListNode

	Tail *SinglyLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateLinkedList(compare CompareFun) *SinglyLinkedList {
	sll := new(SinglyLinkedList)
	sll.size = 0
	sll.compareFunc = compare
	sll.Head = nil
	sll.Tail = nil
	return sll
}

func (sll *SinglyLinkedList) AddAtFront(data interface{}) bool {
	node := new(SinglyLinkedListNode)
	node.Data = data
	node.Next = sll.Head

	sll.Head = node
	sll.size++

	//让tail也指向
	if sll.size == 1 {
		sll.Tail = sll.Head
	}

	return true
}

func (sll *SinglyLinkedList) AddAtEnd(data interface{}) bool {
	node := new(SinglyLinkedListNode)
	node.Data = data
	node.Next = nil

	if sll.size > 0 {
		sll.Tail.Next = node
	}
	sll.Tail = node
	sll.size++

	//让head也指向
	if sll.size == 1 {
		sll.Head = sll.Tail
	}
	return true
}

func (sll SinglyLinkedList) IsEmpty() bool {
	return sll.size == 0
}

func (sll SinglyLinkedList) Size() int {
	return sll.size
}

func (sll *SinglyLinkedList) Search(data interface{}) interface{} {
	for p := sll.Head; p != nil; p = p.Next {
		if sll.compareFunc(p.Data, data) {
			return p.Data
		}
	}

	return nil
}

/**
O(N)
*/
func (sll *SinglyLinkedList) DeleteNode(data interface{}) interface{} {
	if sll.IsEmpty() {
		return false
	}

	var prev *SinglyLinkedListNode = nil

	for p := sll.Head; p != nil; p = p.Next {
		if sll.compareFunc(p.Data, data) {
			if prev != nil {
				prev.Next = p.Next
			}

			if sll.Head == p {
				sll.Head = sll.Head.Next
			}

			if sll.Tail == p {
				sll.Tail = prev
			}

			sll.size--
			return p.Data
		}

		prev = p
	}

	return nil
}

/**
O(N)
*/
func (sll *SinglyLinkedList) DeleteElement(node *SinglyLinkedListNode) bool {
	if sll.IsEmpty() {
		return false
	}

	var prev *SinglyLinkedListNode = nil

	for p := sll.Head; p != nil; p = p.Next {
		if p == node {
			if prev != nil {
				prev.Next = p.Next
			}

			if sll.Head == p {
				sll.Head = sll.Head.Next
			}

			if sll.Tail == p {
				sll.Tail = prev
			}

			sll.size--
			return true
		}

		prev = p
	}

	return false
}
