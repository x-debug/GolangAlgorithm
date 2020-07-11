package datastructs

import (
	"math"
	"math/rand"
)

const MaxLevel = 10

type SkipListNode struct {
	Key   int
	Value interface{}
	Next  []*SkipListNode
}

type SkipList struct {
	Head *SkipListNode

	Tail *SkipListNode

	Size int

	Level int
}

func createNode(key int, value interface{}) *SkipListNode {
	node := new(SkipListNode)
	node.Key = key
	node.Value = value
	node.Next = make([]*SkipListNode, MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		node.Next[i] = nil
	}
	return node
}

func CreateSkipList() *SkipList {
	sl := new(SkipList)
	sl.Level = 0
	sl.Size = 0

	//创建尾哨兵节点
	tail := createNode(math.MaxInt32, nil)
	sl.Tail = tail

	//创建头哨兵节点
	head := createNode(0, nil)
	for i := 0; i < MaxLevel; i++ {
		head.Next[i] = tail
	}
	sl.Head = head
	return sl
}

func (sl SkipList) Search(key int) (bool, interface{}) {
	var p = sl.Head
	for l := sl.Level; l >= 0; l-- {
		for ; p.Next[l].Key < key; {
			p = p.Next[l]
		}
	}

	x := p.Next[0]
	if x.Key == key {
		return true, x.Value
	}
	return false, nil
}

func (sl *SkipList) Insert(key int, value interface{}) bool {
	var p = sl.Head
	var updates []*SkipListNode = make([]*SkipListNode, MaxLevel)

	//通过搜索获得降级节点和要插入的节点
	for l := sl.Level; l >= 0; l-- {
		for ; p.Next[l].Key < key; {
			p = p.Next[l]
		}

		//降级节点
		updates[l] = p
	}

	x := p.Next[0]

	if x.Key == key {
		x.Value = value
		return true //更新节点
	}

	newNode := createNode(key, value)
	level := sl.randomLevel()

	if level > sl.Level {
		//最高级节点
		updates[level] = sl.Head
		sl.Level = level
	}

	//调整指针
	for i := level; i >= 0; i-- {
		x = updates[i]
		newNode.Next[i] = x.Next[i]
		x.Next[i] = newNode
	}

	sl.Size++
	return true
}

func (sl SkipList) randomLevel() int {
	value := rand.Intn(MaxLevel)

	//保持每次只升级1层
	if value > sl.Level {
		value = sl.Level + 1
	}

	return value
}

func (sl *SkipList) Remove(key int) bool {
	p := sl.Head
	var updates []*SkipListNode = make([]*SkipListNode, MaxLevel)

	//查找到需要删除的指针,并且记录降级节点
	for l := sl.Level; l >= 0; l-- {
		for ; p.Next[l].Key < key; {
			p = p.Next[l]
		}

		updates[l] = p
	}

	x := p.Next[0]
	if x.Key != key {
		return false
	}

	//调整指针
	for i := 0; i <= sl.Level; i++ {
		if updates[i].Next[i] != x {
			break
		}

		updates[i].Next[i] = x.Next[i]
	}

	//对层级进行修正
	for ; sl.Level > 0 && sl.Head.Next[sl.Level] == sl.Tail; {
		sl.Level--
	}

	sl.Size--
	return true
}
