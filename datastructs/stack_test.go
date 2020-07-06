package datastructs

import "testing"

func TestCreateStack(t *testing.T) {
	stack := CreateStack()

	if !stack.IsEmpty() {
		t.Error("个数出现了问题")
	}
}

func TestStack_Push(t *testing.T) {
	stack := CreateStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")

	if stack.Size() != 5 {
		t.Error("个数出现了问题")
	}

	stackOrderNodes([]string{"e", "d", "c", "b", "a"}, stack, t)
}

func TestStack_Pop(t *testing.T) {
	stack := CreateStack()

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	stack.Push("e")

	if stack.Size() != 5 {
		t.Error("个数出现了问题")
	}

	stackOrderNodes([]string{"e", "d", "c", "b", "a"}, stack, t)

	stack.Pop()
	stack.Pop()
	stackOrderNodes([]string{"c", "b", "a"}, stack, t)
}

func stackOrderNodes(right []string, s *Stack, t *testing.T) {
	var i int = 0
	for p := s.impl.Head; p != nil; p = p.Next {
		if right[i] != p.Data {
			t.Error("元素次序不对")
		}

		i++
	}
}
