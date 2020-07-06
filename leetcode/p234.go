package leetcode

import (
	"awesomeProject/datastructs"
)

/**
解法1思路:
通过两部走
1).把链表复制到一个固定的数组上
2).通过双指针来判断(分别从前面和后面出发)是否为回文
*/
func P234MethodArray(size int, head *datastructs.SinglyLinkedListNode) bool {
	//链表复制到一个固定的数组上

	ca := make([]int, 0)
	//index := 0
	for p := head; p != nil; p = p.Next {
		//ca[index] = p.Data.(int)
		ca = append(ca, p.Data.(int))
		//index++
	}

	//通过双指针来判断(分别从前面和后面出发)是否为回文
	index := 0
	for ; index < size-1-index; index++ {
		if ca[index] != ca[size-1-index] {
			return false
		}
	}

	return true
}

var gp *datastructs.SinglyLinkedListNode

/**
解法2思路:
通过递归到最后一个的特性,结合每次同时移动节点来做判断
*/
func P234MethodRecursion(size int, head *datastructs.SinglyLinkedListNode) bool {
	gp = head
	return _P234MethodRecursion(head)
}

func _P234MethodRecursion(p *datastructs.SinglyLinkedListNode) bool {
	if p != nil {
		if !_P234MethodRecursion(p.Next) {
			return false
		}

		if p.Data.(int) != gp.Data.(int) {
			return false
		}

		gp = gp.Next
	}

	return true
}

/**
解法3思路:
1.用快慢指针来定位中间节点
2.在定位中间节点的时候开始反转链表
3.再次两个指针从中间和开头开始比较便利
4.恢复反转链表
*/
func P234MethodTwoPointer(size int, head *datastructs.SinglyLinkedListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	var prev *datastructs.SinglyLinkedListNode = nil
	var next *datastructs.SinglyLinkedListNode = nil

	//TODO 快慢指针
	for ; fast != nil && fast.Next != nil; {
		//快指针走两步
		fast = fast.Next.Next

		next = slow.Next

		//反转指针
		slow.Next = prev

		prev = slow
		slow = next
	}

	//如果是奇数和偶数
	if fast != nil && next != nil {
		next = next.Next
	}

	slow = prev
	if slow != nil {
		for p := next; p != nil; p = p.Next {
			if slow.Data.(int) != p.Data.(int) {
				return false
			}
			slow = slow.Next
		}
	}

	return true
}
