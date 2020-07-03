package datastructs

import (
	"fmt"
	"testing"
)

func TestCreateLRUCache(t *testing.T) {
	lru := CreateLRUCache(5)

	lru.Set("hello1", "11")
	lru.Set("hello2", "设计")
	lru.Set("hello3", "100")
	lru.Set("hello4", "产品")
	lru.Set("hello5", "19")
	lru.Set("hello6", "需求")

	//limit to 5
	if lru.Size() != 5 {
		t.Errorf("个数出现了问题, Size():%d", lru.Size())
	}
}

func TestLRUCache_Get(t *testing.T) {
	lru := CreateLRUCache(5)

	lru.Set("hello1", "11")
	lru.Set("hello2", "设计")
	lru.Set("hello3", "100")
	lru.Set("hello4", "产品")
	lru.Set("hello5", "19")
	lru.Set("hello6", "需求")

	//limit to 5
	if lru.Size() != 5 {
		t.Errorf("个数出现了问题, Size():%d", lru.Size())
	}

	value, ok := lru.Get("hello3")
	if ok {
		fmt.Printf("Value:%s\r\n", value)
	}

	lruLoopUp(lru, t)
	fmt.Println(lru.cacheMap)
	if lru.cacheList.head.data.(*LRULinkedElement).key != "hello3" {
		t.Error("Get顺序出现了问题")
	}
}

func TestLRUCache_Set(t *testing.T) {
	lru := CreateLRUCache(5)

	lru.Set("hello1", "11")
	lru.Set("hello2", "设计")
	lru.Set("hello3", "100")
	lru.Set("hello4", "产品")
	lru.Set("hello5", "19")
	lru.Set("hello6", "需求")

	//limit to 5
	if lru.Size() != 5 {
		t.Errorf("个数出现了问题, Size():%d", lru.Size())
	}

	lruLoopUp(lru, t)
	fmt.Println(lru.cacheMap)
	if v, ok := lru.cacheMap["hello1"]; ok {
		t.Errorf("Set出现了问题, %s Len():%d", v.ptr.data, lru.cacheList.Size())
	}
}

func TestLRUCache_Del(t *testing.T) {
	lru := CreateLRUCache(5)

	lru.Set("hello1", "11")
	lru.Set("hello2", "设计")
	lru.Set("hello3", "100")
	lru.Set("hello4", "产品")
	lru.Set("hello5", "19")
	lru.Set("hello6", "需求")

	//limit to 5
	if lru.Size() != 5 {
		t.Errorf("个数出现了问题, Size():%d", lru.Size())
	}

	lruLoopUp(lru, t)
	fmt.Println(lru.cacheMap)

	lru.Del("hello2")
	lru.Del("hello4")

	if lru.Size() != 3 {
		t.Errorf("个数出现了问题, Size():%d", lru.Size())
	}

	lruLoopUp(lru, t)
	fmt.Println(lru.cacheMap)
}

func lruLoopUp(lru *LRUCache, t *testing.T) {
	for p := lru.cacheList.head; p != nil; p = p.next {
		fmt.Println(p.data)
	}
}
