package datastructs

type Queue struct {
	impl *SinglyLinkedList
}

func CreateQueue() *Queue {
	q := new(Queue)
	q.impl = CreateLinkedList(nil)
	return q
}

func (q *Queue) Enqueue(data interface{}) bool {
	q.impl.AddAtEnd(data)
	return true
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.impl.Head.Data
	q.impl.DeleteElement(q.impl.Head)
	return item
}

func (q Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	item := q.impl.Head.Data
	return item
}

func (q Queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q Queue) Size() int {
	return q.impl.Size()
}
