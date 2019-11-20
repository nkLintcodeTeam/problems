// https://leetcode.com/problems/lru-cache/submissions/
package lrucache

import "container/list"

type LRUCache struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:        new(list.List),
		m:        make(map[int]*list.Element, capacity),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.m[key]; ok {
		// after getting this data, this data should be at the front position
		this.l.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if this key alreay is in cache
	if elem, ok := this.m[key]; ok {
		// update value
		elem.Value = Pair{key: key, value: value}
		// move this data to the front
		this.l.MoveToFront(elem)
		return
	}

	// if there are too many keys in the cache, eviction should begin to start
	if this.l.Len() == this.capacity {
		backElem := this.l.Back()
		// we have to delete this data from map and list
		delete(this.m, backElem.Value.(Pair).key)
		this.l.Remove(backElem)
	}

	// the key is not in the cache for now, so just put it in the cache(both list and map)
	newElem := this.l.PushFront(Pair{key: key, value: value})
	this.m[key] = newElem

	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
