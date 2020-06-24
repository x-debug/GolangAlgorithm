package datastructs

import (
	"fmt"
	"math"
	"testing"
)

func comp(a interface{}, b interface{}) bool {
	return a.(int) < b.(int)
}

//测试大数据量下的构建
func TestHeap_Insert(t *testing.T) {
	heap := Create(10000, comp)
	fmt.Println(heap)
	for i := 0; i < 1000; i++ {
		heap.Insert(i + 1)
	}
	fmt.Println(heap)
	checkTreeOk(heap, t)
}

//测试扩容
func TestHeap_Insert2(t *testing.T) {
	heap := Create(10, comp)
	fmt.Println(heap)
	for i := 0; i < 200; i++ {
		heap.Insert(i + 1)
	}
	fmt.Println(heap)
	checkTreeOk(heap, t)
}

func TestHeap_Remove(t *testing.T) {
	heap := Create(100, comp)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(1)
	heap.Insert(7)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(9)
	heap.Insert(8)
	fmt.Println(heap)

	heap.Remove()
	heap.Remove()
	heap.Remove()
	fmt.Println(heap)
	checkTreeOk(heap, t)
}

//测试缩容
func TestHeap_Remove2(t *testing.T) {
	heap := Create(10, comp)
	fmt.Println(heap)
	for i := 0; i < 200; i++ {
		heap.Insert(i + 1)
	}
	fmt.Println(heap)
	checkTreeOk(heap, t)
	for j := 0; j < 150; j++ {
		heap.Remove()
	}
	fmt.Println(heap)
	checkTreeOk(heap, t)
}

func TestHeap_Replace(t *testing.T) {
	heap := Create(100, comp)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(1)
	heap.Insert(7)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(9)
	heap.Insert(8)
	fmt.Println(heap)
	checkTreeOk(heap, t)
	heap.Replace(2, 100)
	checkTreeOk(heap, t)
	fmt.Println(heap)
}

func TestHeap_RemoveAtIndex(t *testing.T) {
	heap := Create(100, comp)
	heap.Insert(4)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(1)
	heap.Insert(7)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(10)
	heap.Insert(9)
	heap.Insert(8)
	fmt.Println(heap)
	checkTreeOk(heap, t)
	heap.RemoveAtIndex(1)
	fmt.Println(heap)
	checkTreeOk(heap, t)
	heap.RemoveAtIndex(4)
	fmt.Println(heap)
	checkTreeOk(heap, t)
	heap.RemoveAtIndex(7)
	fmt.Println(heap)
	checkTreeOk(heap, t)
}

//判断树是否满足最小二叉堆
func checkTreeOk(heap *Heap, t *testing.T) {
	var index uint = heap.len - 1
	for ; index > 0; index-- {
		parent := uint(math.Floor(float64((index - 1) / 2)))
		if parent >= 0 {
			if comp(heap.data[index], heap.data[parent]) {
				t.Error("checkTreeOk error")
			}
		}
	}
}
