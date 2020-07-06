package datastructs

//简单LRU Cache 缓存淘汰算法
type LRUElement struct {
	ptr *DoublyLinkedListNode
}

type LRULinkedElement struct {
	key  string
	data interface{}
}

type LRUCache struct {
	size      int
	maxSize   int
	cacheList *DoublyLinkedList
	cacheMap  map[string]*LRUElement
}

func LRUCacheKeyCompare(a interface{}, b interface{}) bool {
	//ignore
	return true
}

func CreateLRUCache(maxSize int) *LRUCache {
	lru := new(LRUCache)
	lru.maxSize = maxSize
	lru.size = 0
	lru.cacheList = CreateDoublyLinkedList(LRUCacheKeyCompare)
	lru.cacheMap = make(map[string]*LRUElement)
	return lru
}

func (lru LRUCache) Get(key string) (interface{}, bool) {
	ele, ok := lru.cacheMap[key]

	if ok {
		//放到头结点上去
		lru.cacheList.DeleteNode(ele.ptr)
		lru.cacheList.AddFrontAt(ele.ptr)
		return ele.ptr.Data, true
	}

	return nil, false
}

func (lru *LRUCache) Set(key string, value interface{}) bool {
	ele, ok := lru.cacheMap[key]

	if ok {
		lru.cacheList.DeleteNode(ele.ptr)
	} else {
		//淘汰最近没访问过的数据
		lru.size++
		if lru.size > lru.maxSize {
			delete(lru.cacheMap, lru.cacheList.Tail.Data.(*LRULinkedElement).key)
			lru.cacheList.DeleteNode(lru.cacheList.Tail)
			lru.size--
		}
		ele = new(LRUElement)
		listNode := &LRULinkedElement{
			data: value, key: key,
		}
		ele.ptr = CreateDoublyLinkedListNode(listNode)
	}

	item := ele.ptr.Data.(*LRULinkedElement)
	item.data = value
	lru.cacheList.AddFrontAt(ele.ptr)
	ele.ptr = lru.cacheList.Head
	lru.cacheMap[key] = ele
	return true
}

func (lru *LRUCache) Del(key string) bool {
	ele, ok := lru.cacheMap[key]

	if ok {
		lru.cacheList.DeleteNode(ele.ptr)
		delete(lru.cacheMap, key)
		lru.size--
		return true
	}

	return false
}

func (lru LRUCache) Size() int {
	return lru.size
}
