package datastructs

type DoublyLinkedListNode struct {
	prev *DoublyLinkedListNode

	next *DoublyLinkedListNode

	data interface{}
}

type DoublyLinkedList struct {
	head *DoublyLinkedListNode

	tail *DoublyLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateDoublyLinkedList(compare CompareFun) *DoublyLinkedList {
	dll := new(DoublyLinkedList)
	dll.head = nil
	dll.tail = nil
	dll.size = 0
	dll.compareFunc = compare
	return dll
}

func (dll *DoublyLinkedList) AddFront(data interface{}) bool {
	node := new(DoublyLinkedListNode)
	node.data = data

	return dll.AddFrontAt(node)
}

func (dll *DoublyLinkedList) AddFrontAt(node *DoublyLinkedListNode) bool {
	node.prev = nil

	if dll.size == 0 { //it is empty
		dll.head = node
		dll.tail = node
		node.next = nil //pointer to self
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddEndAt(node *DoublyLinkedListNode) bool {
	node.next = nil

	if dll.size == 0 { //it is empty
		dll.head = node
		dll.tail = node
		node.prev = nil //pointer to self
	} else {
		node.prev = dll.tail
		dll.tail.next = node
		dll.tail = node
	}

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddEnd(data interface{}) bool {
	node := new(DoublyLinkedListNode)
	node.data = data

	return dll.AddEndAt(node)
}

func (dll *DoublyLinkedList) AddBefore(at *DoublyLinkedListNode, data interface{}) bool {
	if at == nil {
		return false
	}

	node := new(DoublyLinkedListNode)
	node.data = data

	node.prev = at.prev
	node.next = at

	if at.prev != nil {
		at.prev.next = node
	} else {
		dll.head = node
	}
	at.prev = node

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddAfter(at *DoublyLinkedListNode, data interface{}) bool {
	if at == nil {
		return false
	}

	node := new(DoublyLinkedListNode)
	node.data = data

	node.next = at.next
	node.prev = at

	if at.next != nil {
		at.next.prev = node
	} else {
		dll.tail = node
	}
	at.next = node

	dll.size++
	return true
}

func (dll DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

func (dll DoublyLinkedList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList) DeleteNode(at *DoublyLinkedListNode) bool {
	if dll.IsEmpty() {
		return false
	}

	if at == nil {
		return false
	}

	if at.prev != nil {
		at.prev.next = at.next
	} else { //pointer to head
		dll.head = at.next
	}

	if at.next != nil {
		at.next.prev = at.prev
	} else { //pointer to next
		dll.tail = at.prev
	}

	dll.size--
	return true
}

func (dll *DoublyLinkedList) SearchForward(data interface{}) *DoublyLinkedListNode {
	if dll.IsEmpty() {
		return nil
	}

	for p := dll.tail; p != nil; p = p.prev {
		if dll.compareFunc(p.data, data) {
			return p
		}
	}

	return nil
}

func (dll *DoublyLinkedList) SearchBackward(data interface{}) *DoublyLinkedListNode {
	if dll.IsEmpty() {
		return nil
	}

	for p := dll.head; p != nil; p = p.next {
		if dll.compareFunc(p.data, data) {
			return p
		}
	}

	return nil
}
