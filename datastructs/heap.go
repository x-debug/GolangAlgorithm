package datastructs

import (
	"fmt"
	"math"
)

type CompareFun func(a interface{}, b interface{}) bool

/**
最小堆
*/
type Heap struct {
	cap  uint
	len  uint
	data []interface{}
	comp CompareFun
}

/**
向上修复 O(logN)
*/
func (heap *Heap) ShiftUp(index uint) {
	if index > 0 && index < heap.len {
		parent := uint(math.Floor(float64((index - 1) / 2)))
		if heap.comp(heap.data[index], heap.data[parent]) {
			heap.data[index], heap.data[parent] = heap.data[parent], heap.data[index]
			//fmt.Printf("child<->parent:%d,%d\r\n", index, parent)
		}

		heap.ShiftUp(parent)
	}
}

/**
向下修复 O(logN)
*/
func (heap *Heap) ShiftDown(index uint) {
	if index >= 0 && index < heap.len {
		left := 2*index + 1
		right := 2*index + 2
		if left < heap.len {
			if right < heap.len {
				if heap.comp(heap.data[right] , heap.data[left]) {
					left = right
				}
			}
			if heap.comp(heap.data[left] , heap.data[index]) {
				heap.data[left], heap.data[index] = heap.data[index], heap.data[left]
				//fmt.Printf("child<->parent:%d,%d\r\n", left, index)
			}

			heap.ShiftDown(left)
		}
	}
}

func Create(cap uint, comp CompareFun) *Heap {
	heap := new(Heap)
	heap.len = 0
	heap.cap = cap
	heap.data = make([]interface{}, heap.cap)
	heap.comp = comp
	return heap
}

/**
	O(logN)
 */
func (heap *Heap) Insert(element int) bool {
	if heap.len >= heap.cap {
		//扩容
		newBuffer := make([]interface{}, heap.cap*2)
		copy(newBuffer, heap.data)
		heap.data = newBuffer
		heap.cap = heap.cap * 2
		fmt.Printf("heap 进行了一次扩容, 当然容量:%d\r\n", heap.cap)
	}

	heap.data[heap.len] = element
	heap.len++
	heap.ShiftUp(heap.len - 1)
	return true
}

/**
	O(logN)
*/
func (heap *Heap) Remove() interface{} {
	if heap.isEmpty() {
		panic("heap is empty")
	}

	if heap.len < heap.cap/2 {
		//缩容
		newBuffer := make([]interface{}, heap.cap/2)
		copy(newBuffer, heap.data)
		heap.data = newBuffer
		heap.cap = heap.cap / 2
		fmt.Printf("heap 进行了一次缩容, 当然容量:%d\r\n", heap.cap)
	}

	value := heap.data[0]
	heap.data[0] = heap.data[heap.len-1]
	heap.data[heap.len-1] = 0
	heap.len--
	heap.ShiftDown(0)
	return value
}

func (heap Heap) isEmpty() bool {
	return heap.len == 0
}

/**
	O(N)
*/
func (heap *Heap) RemoveAtIndex(index uint) interface{} {
	if heap.isEmpty() {
		panic("heap is empty")
	}

	if index < 0 || index >= heap.len {
		panic("range out")
	}

	value := heap.data[index]

	for i := index; i < heap.len; i++ {
		if i >= 0 {
			heap.data[i] = heap.data[i+1]
		}
	}
	heap.len--
	heap.ShiftDown(0)

	return value
}

/**
	O(logN)
*/
func (heap *Heap) Replace(index uint, element int) bool {
	if heap.isEmpty() {
		return false
	}

	if index < 0 || index >= heap.len {
		return false
	}

	heap.data[index] = element
	heap.ShiftUp(index)
	heap.ShiftDown(index)

	return true
}
