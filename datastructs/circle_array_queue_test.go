package datastructs

import (
	"fmt"
	"testing"
)

func TestCreateCircleArrayQueue(t *testing.T) {
	caq := CreateCircleArrayQueue(5)
	caq.Enqueue(1)
	caq.Enqueue(2)
	caq.Enqueue(3)
	//caq.Enqueue(4)
	caqLoopUp(caq)
}

func TestCircleArrayQueue_Enqueue(t *testing.T) {
	caq := CreateCircleArrayQueue(5)
	caq.Enqueue(1)
	caq.Enqueue(2)
	caq.Enqueue(3)
	//caq.Enqueue(4)
	caqLoopUp(caq)
	if caq.Size != 3 {
		t.Error("TestCircleArrayQueue_Enqueue")
	}
}

func TestCircleArrayQueue_Dequeue(t *testing.T) {
	caq := CreateCircleArrayQueue(5)
	caq.Enqueue(1)
	caq.Enqueue(2)
	caq.Enqueue(3)
	//caq.Enqueue(4)
	caqLoopUp(caq)

	caq.Dequeue()
	caq.Dequeue()
	caq.Dequeue()
	//caq.Dequeue()
	caqLoopUp(caq)

	if caq.Size != 0 {
		t.Error("TestCircleArrayQueue_Enqueue")
	}

	caq.Enqueue(1)
	caq.Enqueue(2)
	caq.Enqueue(3)
	caq.Enqueue(3)
	caq.Enqueue(3)
	caqLoopUp(caq)
	if caq.Size != 5 {
		t.Error("TestCircleArrayQueue_Enqueue")
	}
}

func caqLoopUp(queue *CircleArrayQueue) {
	fmt.Println(queue.Data)
}
