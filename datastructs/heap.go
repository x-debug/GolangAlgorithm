package datastructs

import (
	"fmt"
	"math"
)

/**
最小堆
*/
type Heap struct {
	cap  uint
	len  uint
	Data []interface{}
	comp CompareFun
}

/**
向上修复 O(logN)
*/
func (heap *Heap) ShiftUp(index uint) {
	for {
		if index > 0 && index < heap.len {
			parent := uint(math.Floor(float64((index - 1) / 2)))
			if heap.comp(heap.Data[index], heap.Data[parent]) {
				heap.Data[index], heap.Data[parent] = heap.Data[parent], heap.Data[index]
				//fmt.Printf("child<->parent:%d,%d\r\n", index, parent)
			}
			index = parent
		} else {
			break
		}
	}
}

/**
向下修复 O(logN)
*/
func (heap *Heap) ShiftDown(index uint) {
	for {
		if index >= 0 && index < heap.len {
			left := 2*index + 1
			right := 2*index + 2
			if left < heap.len {
				if right < heap.len {
					if heap.comp(heap.Data[right], heap.Data[left]) {
						left = right
					}
				}
				if heap.comp(heap.Data[left], heap.Data[index]) {
					heap.Data[left], heap.Data[index] = heap.Data[index], heap.Data[left]
					//fmt.Printf("child<->parent:%d,%d\r\n", left, index)
				}

				index = left
			} else {
				break
			}
		} else {
			break
		}
	}
}

func CreateHeap(cap uint, comp CompareFun) *Heap {
	heap := new(Heap)
	heap.len = 0
	heap.cap = cap
	heap.Data = make([]interface{}, heap.cap)
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
		copy(newBuffer, heap.Data)
		heap.Data = newBuffer
		heap.cap = heap.cap * 2
		fmt.Printf("heap 进行了一次扩容, 当然容量:%d\r\n", heap.cap)
	}

	heap.Data[heap.len] = element
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
		copy(newBuffer, heap.Data)
		heap.Data = newBuffer
		heap.cap = heap.cap / 2
		fmt.Printf("heap 进行了一次缩容, 当然容量:%d\r\n", heap.cap)
	}

	value := heap.Data[0]
	heap.Data[0] = heap.Data[heap.len-1]
	heap.Data[heap.len-1] = 0
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

	value := heap.Data[index]

	for i := index; i < heap.len; i++ {
		if i >= 0 {
			heap.Data[i] = heap.Data[i+1]
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

	heap.Data[index] = element
	heap.ShiftUp(index)
	heap.ShiftDown(index)

	return true
}
