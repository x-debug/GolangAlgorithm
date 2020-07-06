package datastructs

type CircularLinkedListNode struct {
	Next *CircularLinkedListNode

	Data interface{}
}

type CircularLinkedList struct {
	Head *CircularLinkedListNode

	Tail *CircularLinkedListNode

	compareFunc CompareFun

	size int
}

func CreateCircularLinkedList(compare CompareFun) *CircularLinkedList {
	cll := new(CircularLinkedList)
	cll.size = 0
	cll.compareFunc = compare
	cll.Head = nil
	return cll
}

func (cll *CircularLinkedList) AddAtFront(data interface{}) bool {
	node := new(CircularLinkedListNode)
	node.Data = data

	if cll.size == 0 { //it is empty
		cll.Head = node
		cll.Tail = node
		node.Next = node //pointer to self
	} else {
		node.Next = cll.Head
		cll.Head = node
		cll.Tail.Next = cll.Head
	}

	cll.size++
	return true
}

func (cll *CircularLinkedList) AddAtEnd(data interface{}) bool {
	node := new(CircularLinkedListNode)
	node.Data = data

	if cll.size == 0 { //it is empty
		cll.Head = node
		cll.Tail = node
		node.Next = node //pointer to self
	} else {
		cll.Tail.Next = node
		cll.Tail = node
		node.Next = cll.Head
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
	for p := cll.Head; index < cll.size; p = p.Next {
		if cll.compareFunc(p.Data, data) {
			return p.Data
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

	for p := cll.Head; index < cll.size; p = p.Next {
		if cll.compareFunc(p.Data, data) {
			if prev != nil {
				prev.Next = p.Next
			}

			if cll.Head == p {
				cll.Head = cll.Head.Next
			}

			if cll.Tail == p {
				cll.Tail = prev
			}

			cll.size--
			return p.Data
		}

		prev = p
		index++
	}

	return nil
}
