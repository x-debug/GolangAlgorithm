package datastructs

type SinglyLinkedListNode struct {
	next *SinglyLinkedListNode

	data interface{}
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode

	tail *SinglyLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateLinkedList(compare CompareFun) *SinglyLinkedList {
	sll := new(SinglyLinkedList)
	sll.size = 0
	sll.compareFunc = compare
	sll.head = nil
	sll.tail = nil
	return sll
}

func (sll *SinglyLinkedList) AddAtFront(data interface{}) bool {
	node := new(SinglyLinkedListNode)
	node.data = data
	node.next = sll.head

	sll.head = node
	sll.size++

	//让tail也指向
	if sll.size == 1 {
		sll.tail = sll.head
	}

	return true
}

func (sll *SinglyLinkedList) AddAtEnd(data interface{}) bool {
	node := new(SinglyLinkedListNode)
	node.data = data
	node.next = nil

	if sll.size > 0 {
		sll.tail.next = node
	}
	sll.tail = node
	sll.size++

	//让head也指向
	if sll.size == 1 {
		sll.head = sll.tail
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
	for p := sll.head; p != nil; p = p.next {
		if sll.compareFunc(p.data, data) {
			return p.data
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

	for p := sll.head; p != nil; p = p.next {
		if sll.compareFunc(p.data, data) {
			if prev != nil {
				prev.next = p.next
			}

			if sll.head == p {
				sll.head = sll.head.next
			}

			if sll.tail == p {
				sll.tail = prev
			}

			sll.size--
			return p.data
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

	for p := sll.head; p != nil; p = p.next {
		if p == node {
			if prev != nil {
				prev.next = p.next
			}

			if sll.head == p {
				sll.head = sll.head.next
			}

			if sll.tail == p {
				sll.tail = prev
			}

			sll.size--
			return true
		}

		prev = p
	}

	return false
}
