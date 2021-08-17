package main

import "container/list"

type LRUCache struct {
	capacity int
	mem      map[int](*list.Element)
	cache    list.List
}

func Constructor(capacity int) LRUCache {
	newCache := LRUCache{}
	newCache.capacity = capacity
	newCache.mem = make(map[int]*list.Element, capacity)
	newCache.cache = list.List{}
	return newCache
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.mem[key]; ok {

		this.cache.MoveToFront(v)
		return v.Value.([2]int)[1]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.mem[key]; ok {

		this.cache.MoveToFront(v)
		v.Value = [2]int{key, value}
	} else {
		if this.cache.Len() == this.capacity {
			delete(this.mem, this.cache.Back().Value.([2]int)[0])
			this.cache.Remove(this.cache.Back())
		}
		this.cache.PushFront([2]int{key, value})
		this.mem[key] = this.cache.Front()
	}
}

func main() {

}
