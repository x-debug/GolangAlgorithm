package datastructs

import "testing"

func TestCreateQueue(t *testing.T) {
	queue := CreateQueue()

	if !queue.IsEmpty() {
		t.Error("个数出现了问题")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	queue := CreateQueue()

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")

	if queue.Size() != 5 {
		t.Error("个数出现了问题")
	}

	queueOrderNodes([]string{"a", "b", "c", "d", "e"}, queue, t)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := CreateQueue()

	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queue.Enqueue("e")

	if queue.Size() != 5 {
		t.Error("个数出现了问题")
	}

	queueOrderNodes([]string{"a", "b", "c", "d", "e"}, queue, t)

	queue.Dequeue()//b c d e
	queue.Enqueue("fuck")//b c d e fuck
	queue.Enqueue("baby")//b c d e fuck baby
	queue.Dequeue()//c d e fuck baby
	queue.Dequeue()//d e fuck baby
	if queue.Size() != 4 {
		t.Error("个数出现了问题")
	}

	queueOrderNodes([]string{"d", "e", "fuck", "baby"}, queue, t)
}

func queueOrderNodes(right []string, s *Queue, t *testing.T) {
	var i int = 0
	for p := s.impl.head; p != nil; p = p.next {
		if right[i] != p.data {
			t.Error("元素次序不对")
			break
		}

		i++
	}
}
