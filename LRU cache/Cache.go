package main

import "container/list"

// type LRUCache struct {
// 	mp   map[int]*ListNode
// 	head *ListNode
// 	tail *ListNode
// 	cap  int
// }

// func newLRUCache(capacity int) *LRUCache {
// 	LRUCache := &LRUCache{
// 		mp:   make(map[int]*ListNode),
// 		head: newLisNode(),
// 		tail: newLisNode(),
// 		cap:  capacity,
// 	}
// 	LRUCache.head.next = LRUCache.tail
// 	LRUCache.tail.prev = LRUCache.head
// 	return LRUCache
// }

// func (cache *LRUCache) get(key int) int {
// 	if node, ok := cache.mp[key]; ok {
// 		node.next = cache.tail
// 		node.prev = cache.tail.prev
// 		cache.tail.prev.next = node
// 		return node.val
// 	}
// 	return -1
// }

// func (cache *LRUCache) put(key int, val int) {
// 	if node, ok := cache.mp[key]; ok {
// 		node.val = val
// 		cache.tail.prev.next = node
// 		cache.tail.prev.prev = node.prev
// 		node.next = cache.tail
// 		node.prev = cache.tail.prev
// 		cache.tail.prev = node

// 	} else if len(cache.mp) <= cache.cap {
// 		node = newLisNode()
// 		node.key = key
// 		node.val = val
// 		node.next = cache.tail
// 		node.prev = cache.tail.prev
// 		cache.tail.prev.next = node
// 		cache.tail.prev = node
// 		cache.mp[key] = node
// 	} else {
// 		node = newLisNode()
// 		node.key = key
// 		node.val = val
// 		node.next = cache.tail
// 		node.prev = cache.tail.prev
// 		cache.tail.prev.next = node
// 		cache.tail.prev = node
// 		delete(cache.mp, cache.head.next.key)
// 		temp := cache.head.next
// 		cache.head.next = temp.next
// 		temp = nil
// 		cache.mp[key] = node
// 	}
// }

type LRUCache struct {
	Cap int
	L   *list.List
	M   map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap: capacity,
		L:   list.New(),
		M:   make(map[int]*list.Element, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.M[key]; ok {
		this.L.MoveToBack(node)
		return node.Value.(*Node).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.M[key]; ok {
		node.Value.(*Node).Value = value
		this.L.MoveToBack(node)
	} else {
		newNode := &Node{Key: key, Value: value}
		if len(this.M) < this.Cap {
			ptr := this.L.PushBack(newNode)
			this.M[key] = ptr
		} else {
			frontNode := this.L.Front()
			this.L.Remove(frontNode)
			delete(this.M, frontNode.Value.(*Node).Key)
			ptr := this.L.PushBack(newNode)
			this.M[key] = ptr
		}

	}
}
