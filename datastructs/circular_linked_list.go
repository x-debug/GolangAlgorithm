package datastructs

type CircularLinkedListNode struct {
	next *CircularLinkedListNode

	data interface{}
}

type CircularLinkedList struct {
	head *CircularLinkedListNode

	tail *CircularLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateCircularLinkedList(compare CompareFun) *CircularLinkedList {
	cll := new(CircularLinkedList)
	cll.size = 0
	cll.compareFunc = compare
	cll.head = nil
	return cll
}

func (cll *CircularLinkedList) AddAtFront(data interface{}) bool {
	node := new(CircularLinkedListNode)
	node.data = data

	if cll.size == 0 { //it is empty
		cll.head = node
		cll.tail = node
		node.next = node //pointer to self
	} else {
		node.next = cll.head
		cll.head = node
		cll.tail.next = cll.head
	}

	cll.size++
	return true
}

func (cll *CircularLinkedList) AddAtEnd(data interface{}) bool {
	node := new(CircularLinkedListNode)
	node.data = data

	if cll.size == 0 { //it is empty
		cll.head = node
		cll.tail = node
		node.next = node //pointer to self
	} else {
		cll.tail.next = node
		cll.tail = node
		node.next = cll.head
	}

	cll.size++
	return true
}

func (cll CircularLinkedList) IsEmpty() bool {
	return cll.size == 0
}

func (cll CircularLinkedList) Size() int {
	return cll.size
}

func (cll *CircularLinkedList) Search(data interface{}) interface{} {
	if cll.IsEmpty() {
		return nil
	}

	index := 0
	for p := cll.head; index < cll.size; p = p.next {
		if cll.compareFunc(p.data, data) {
			return p.data
		}

		index++
	}

	return nil
}

func (cll *CircularLinkedList) DeleteNode(data interface{}) interface{} {
	if cll.IsEmpty() {
		return nil
	}

	index := 0
	var prev *CircularLinkedListNode = nil

	for p := cll.head; index < cll.size; p = p.next {
		if cll.compareFunc(p.data, data) {
			if prev != nil {
				prev.next = p.next
			}

			if cll.head == p {
				cll.head = cll.head.next
			}

			if cll.tail == p {
				cll.tail = prev
			}

			cll.size--
			return p.data
		}

		prev = p
		index++
	}

	return nil
}
