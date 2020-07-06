package datastructs

type Stack struct {
	impl *SinglyLinkedList
}

func CreateStack() *Stack {
	s := new(Stack)
	s.impl = CreateLinkedList(nil)
	return s
}

func (s *Stack) Push(data interface{}) bool {
	s.impl.AddAtFront(data)
	return true
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.impl.Head.Data
	s.impl.DeleteElement(s.impl.Head)

	return data
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return s.impl.Size()
}
