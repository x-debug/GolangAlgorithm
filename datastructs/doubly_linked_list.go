package datastructs

type DoublyLinkedListNode struct {
	Prev *DoublyLinkedListNode

	Next *DoublyLinkedListNode

	Data interface{}
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode

	Tail *DoublyLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateDoublyLinkedList(compare CompareFun) *DoublyLinkedList {
	dll := new(DoublyLinkedList)
	dll.Head = nil
	dll.Tail = nil
	dll.size = 0
	dll.compareFunc = compare
	return dll
}

func CreateDoublyLinkedListNode(data interface{}) *DoublyLinkedListNode {
	node := new(DoublyLinkedListNode)
	node.Data = data

	return node
}

func (dll *DoublyLinkedList) AddFront(data interface{}) bool {
	node := new(DoublyLinkedListNode)
	node.Data = data

	return dll.AddFrontAt(node)
}

func (dll *DoublyLinkedList) AddFrontAt(node *DoublyLinkedListNode) bool {
	node.Prev = nil

	if dll.size == 0 { //it is empty
		dll.Head = node
		dll.Tail = node
		node.Next = nil //pointer to self
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		dll.Head = node
	}

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddEndAt(node *DoublyLinkedListNode) bool {
	node.Next = nil

	if dll.size == 0 { //it is empty
		dll.Head = node
		dll.Tail = node
		node.Prev = nil //pointer to self
	} else {
		node.Prev = dll.Tail
		dll.Tail.Next = node
		dll.Tail = node
	}

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddEnd(data interface{}) bool {
	node := new(DoublyLinkedListNode)
	node.Data = data

	return dll.AddEndAt(node)
}

func (dll *DoublyLinkedList) AddBefore(at *DoublyLinkedListNode, data interface{}) bool {
	if at == nil {
		return false
	}

	node := new(DoublyLinkedListNode)
	node.Data = data

	node.Prev = at.Prev
	node.Next = at

	if at.Prev != nil {
		at.Prev.Next = node
	} else {
		dll.Head = node
	}
	at.Prev = node

	dll.size++
	return true
}

func (dll *DoublyLinkedList) AddAfter(at *DoublyLinkedListNode, data interface{}) bool {
	if at == nil {
		return false
	}

	node := new(DoublyLinkedListNode)
	node.Data = data

	node.Next = at.Next
	node.Prev = at

	if at.Next != nil {
		at.Next.Prev = node
	} else {
		dll.Tail = node
	}
	at.Next = node

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

	if at.Prev != nil {
		at.Prev.Next = at.Next
	} else { //pointer to Head
		dll.Head = at.Next
	}

	if at.Next != nil {
		at.Next.Prev = at.Prev
	} else { //pointer to Next
		dll.Tail = at.Prev
	}

	dll.size--
	return true
}

func (dll *DoublyLinkedList) SearchForward(data interface{}) *DoublyLinkedListNode {
	if dll.IsEmpty() {
		return nil
	}

	for p := dll.Tail; p != nil; p = p.Prev {
		if dll.compareFunc(p.Data, data) {
			return p
		}
	}

	return nil
}

func (dll *DoublyLinkedList) SearchBackward(data interface{}) *DoublyLinkedListNode {
	if dll.IsEmpty() {
		return nil
	}

	for p := dll.Head; p != nil; p = p.Next {
		if dll.compareFunc(p.Data, data) {
			return p
		}
	}

	return nil
}
