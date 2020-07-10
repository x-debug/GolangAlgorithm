package datastructs

type CircleArrayQueue struct {
	MaxSize int

	Size int

	Head int

	Tail int

	Data []interface{}
}

func CreateCircleArrayQueue(maxSize int) *CircleArrayQueue {
	caq := new(CircleArrayQueue)
	caq.MaxSize = maxSize + 1
	caq.Size = 0
	caq.Head = 0
	caq.Tail = 0
	caq.Data = make([]interface{}, caq.MaxSize)
	return caq
}

func (caq *CircleArrayQueue) Enqueue(data interface{}) bool {
	if (caq.Tail+1)%caq.MaxSize == caq.Head {
		panic("Queue overflow")
	}

	caq.Tail = (caq.Tail + 1) % caq.MaxSize
	caq.Data[caq.Tail] = data
	caq.Size++
	return true
}

func (caq *CircleArrayQueue) Dequeue() interface{} {
	if caq.Head == caq.Tail {
		panic("Queue empty")
	}

	caq.Head = (caq.Head + 1) % caq.MaxSize
	v := caq.Data[caq.Head]
	caq.Data[caq.Head] = nil//clear
	caq.Size--
	return v
}
