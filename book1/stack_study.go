package book1

import (
	"awesomeProject/datastructs"
)

/**
时间复杂度 O(N)
用两个操作数栈和操作符栈来计算四则运算的结果

核心想法:通过栈的特殊结构压入栈中，然后直到遇到操作符号小于前面的操作符号才开始计算
*/
func ScanOperateNum(rawNum string) *datastructs.SinglyLinkedList {
	numStream := datastructs.CreateLinkedList(nil)
	var t int = 0
	for _, v := range rawNum {
		if !isOperate((int)(v)) {
			t = 10*t + int(v) - 48
		} else {
			numStream.AddAtEnd(t)
			t = 0
			numStream.AddAtEnd((int)(v))
		}
	}
	numStream.AddAtEnd(t)
	return numStream
}

func isOperate(ch int) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isLess(ch int) int {
	if ch == '+' || ch == '-' {
		return 0
	}

	if ch == '*' || ch == '/' {
		return 1
	}

	return 0
}

/**
时间复杂度 O(N)
 */
func ComputeNum(input *datastructs.SinglyLinkedList) int {
	numStack := datastructs.CreateStack()
	operStack := datastructs.CreateStack()

	for p := input.Head; p != nil; p = p.Next {
		ch := p.Data.(int)
		if !isOperate(ch) {
			numStack.Push(ch)
		} else {
			for {
				v := operStack.Peek()

				if v != nil {
					if isLess(ch) <= isLess(v.(int)) { //遇到优先度底的运算符,开始前面计算
						n1 := numStack.Pop().(int)
						n2 := numStack.Pop().(int)
						o := operStack.Pop().(int)
						n := computeTwo(o, n1, n2)
						numStack.Push(n)
					} else {
						operStack.Push(ch)
						break
					}
				} else {
					operStack.Push(ch)
					break
				}
			}
		}
	}

	n1 := numStack.Pop().(int)
	n2 := numStack.Pop().(int)
	o := operStack.Pop().(int)
	return computeTwo(o, n1, n2)
}

func computeTwo(o int, n1 int, n2 int) int {
	var n int = 0
	switch o {
	case '+':
		n = n1 + n2
	case '-':
		n = n2 - n1
	case '*':
		n = n1 * n2
	case '/':
		n = n2 / n1

	}

	return n
}
