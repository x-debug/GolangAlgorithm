package datastructs

import "fmt"

const MsCap = 10

type EventFun func(ms *MultiSet, data interface{}, index uint)

type MultiSet struct {
	used        uint
	cap         uint
	last        uint
	data        []interface{}
	onInsert    EventFun
	onRemove    EventFun
	compareFunc CompareFun
}

func CreateMS(compareFun CompareFun, onInsert EventFun, onRemove EventFun) *MultiSet {
	ms := new(MultiSet)
	ms.used = 0
	ms.cap = MsCap
	ms.last = 0
	ms.data = make([]interface{}, ms.cap)
	ms.onInsert = onInsert
	ms.onRemove = onRemove
	ms.compareFunc = compareFun
	return ms
}

func (ms *MultiSet) delete(index uint) bool {
	if ms.used <= 0 || index >= ms.used {
		return false
	}

	item := ms.data[index]
	//把最后一个元素复制到删除的那个元素位置
	ms.data[index] = ms.data[ms.used-1]
	ms.used--

	if ms.onRemove != nil {
		ms.onRemove(ms, item, index)
	}
	return true
}

func (ms *MultiSet) grow() {
	ms.cap = ms.cap * 2
	newBuffer := make([]interface{}, ms.cap)
	copy(newBuffer, ms.data)
	ms.data = newBuffer
}

/**
O(1)
*/
func (ms *MultiSet) Append(data interface{}) bool {
	if ms.used == ms.cap {
		ms.grow() //进行扩容
		fmt.Printf("MS 进行了一次扩容,扩容后的容量为:%d", ms.cap)
	}

	if ms.used == ms.cap {
		return false
	}

	ms.data[ms.used] = data
	ms.used++
	if ms.onInsert != nil {
		ms.onInsert(ms, data, ms.used-1)
	}
	return true
}

func (ms *MultiSet) Clear() {
	for {
		if !ms.delete(0) {
			break
		}
	}

	ms.cap = 0
	ms.last = 0
}

/**
O(N)
*/
func (ms *MultiSet) Remove(data interface{}) bool {
	var i uint = 0
	for ; i < ms.used; i++ {
		if ms.compareFunc(ms.data[i], data) {
			ms.delete(i)
			return true
		}
	}

	return false
}

/**
O(N)
*/
func (ms MultiSet) Contains(data interface{}) bool {
	var i uint = 0
	for ; i < ms.used; i++ {
		if ms.compareFunc(ms.data[i], data) {
			return true
		}
	}

	return false
}

/**
O(1)
*/
func (ms *MultiSet) Take() interface{} {
	if ms.isEmpty() {
		return nil
	}

	ms.last = ms.last % ms.used
	item := ms.data[ms.last]
	ms.delete(ms.last)
	ms.last++
	return item
}

func (ms MultiSet) isEmpty() bool {
	return ms.used == 0
}
