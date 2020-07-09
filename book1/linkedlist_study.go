package book1

import "awesomeProject/datastructs"

/**
O(N)
反转链表思路:
在迭代的同时，进行指针的反向逆转
*/
func ReverseLinkedList(linked *datastructs.SinglyLinkedList) {
	var prev *datastructs.SinglyLinkedListNode = nil
	var k *datastructs.SinglyLinkedListNode = nil

	for p := linked.Head; p != nil; {
		k = p
		p = p.Next
		k.Next = prev
		prev = k
	}

	linked.Head = k
}

/**
O(N) = O(K1) + O(K2)
合并链表思路:
因为是两个链表是有序的，所有只要迭代比较两个链表，然后选取较小的元素放入新链表中，并且把剩下的元素放入新链表中返回即可
*/
func JoinTwoLinkedList1(A *datastructs.SinglyLinkedList, B *datastructs.SinglyLinkedList) *datastructs.SinglyLinkedList {
	head1 := A.Head
	head2 := B.Head

	result := datastructs.CreateLinkedList(nil)

	for {
		if head1.Data.(int) < head2.Data.(int) {
			result.AddAtEnd(head1.Data)
			head1 = head1.Next
		} else {
			result.AddAtEnd(head2.Data)
			head2 = head2.Next
		}

		if head1 == nil || head2 == nil {
			if head1 != nil {
				for p := head1; p != nil; p = p.Next {
					result.AddAtEnd(p.Data)
				}
			}

			if head2 != nil {
				for p := head2; p != nil; p = p.Next {
					result.AddAtEnd(p.Data)
				}
			}

			break
		}
	}

	return result
}

/**

这种方法会改变A,B原来的链表
1) 先在A中找到比B大的要插入的K点, 插入后弹出B链表节点
2) 递归执行1)过程,A=K
3) B为空,直接结束返回A就可,如果A为空,把B添加到A末尾返回即可
*/
func JoinTwoLinkedList2(A *datastructs.SinglyLinkedList, B *datastructs.SinglyLinkedList) *datastructs.SinglyLinkedList {
	return nil
}

/**
利用快慢指针法可以轻松定位中间节点
*/
func GetLinkedMiddleNode(ll *datastructs.SinglyLinkedList) *datastructs.SinglyLinkedListNode {
	fast := ll.Head
	slow := ll.Head
	var prev *datastructs.SinglyLinkedListNode
	var p *datastructs.SinglyLinkedListNode

	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		p = slow.Next
		slow.Next = prev
		prev = slow
		slow = p
	}

	return slow
}

/**
利用尾递归来从后面倒数K个节点
*/
func DeleteLinkedNode(ll *datastructs.SinglyLinkedList, k int) {
	gs = datastructs.CreateStack()
	iterLinkedNode(ll, ll.Head, k)
}

/**
 采用HashMap方法很好理解,时间复杂度O(N),空间复杂度O(N)
这里采用快慢指针思路:

当快指针首先到链表末尾，则链表不存在Circle,反之快指针一定会追上慢指针
*/
func CheckCircleLinked(ll *datastructs.SinglyLinkedList) bool {
	if ll.Head == nil || ll.Head.Next == nil {
		return false
	}

	fast := ll.Head
	slow := ll.Head

	for ; fast != nil && fast.Next != nil; {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

var gi int = 0
var gs *datastructs.Stack

func iterLinkedNode(ll *datastructs.SinglyLinkedList, node *datastructs.SinglyLinkedListNode, k int) {
	if node != nil {
		prev := gs.Pop()
		gs.Push(node)
		iterLinkedNode(ll, node.Next, k)
		gi++
		if gi == k {
			if prev != nil {
				//delete node
				tmp := prev.(*datastructs.SinglyLinkedListNode)
				tmp.Next = node.Next
			} else {
				ll.Head = node.Next
			}
		}
	}
}
