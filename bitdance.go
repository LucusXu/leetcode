package main

import "fmt"

type LinkNode struct {
	key int
	val int
	prev *LinkNode
	next *LinkNode
}

type LRU struct {
	header *LinkNode
	tail *LinkNode
	cache map[int]*LinkNode
	cap int
}

func Construct(capacity int) LRU {
	lru := LRU{
		header : &LinkNode{},
		tail : &LinkNode{},
		cap : capacity,
		cache : make(map[int]*LinkNode, capacity),
	}
	lru.header.next = lru.tail
	lru.tail.prev = lru.header
	return lru
}

func (this *LRU) get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.remove(node)
		this.putToHead(node)
		return node.val
	}
}

func (this *LRU) put(key, value int) {
	if node, ok := this.cache[key]; ok {
		this.cache[key].val = value
		this.remove(node)
		this.putToHead(node)
	} else {
		if this.cap <= len(this.cache) {
			delNode := this.tail.prev
			delete(this.cache, delNode.key)
			this.remove(delNode)
		}
		newNode := LinkNode{
			key: key,
			val: value,
		}
		this.putToHead(&newNode)
		this.cache[key] = &newNode
	}
}

func (this *LRU) remove(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRU) putToHead(node *LinkNode) {
	tmp := this.header.next
	node.next = tmp
	tmp.prev = node
	node.prev = this.header
	this.header.next = node
}


func main() {
	lru := Construct(3)
	lru.put(1,1)
	lru.put(2,2)
	lru.put(3,3)
	lru.put(4,4)
	res := lru.get(0)
	fmt.Println("%s\n", res)
	res = lru.get(1)
	fmt.Println("%v\n", res)
	res = lru.get(2)
	fmt.Println("%v\n", res)
	res = lru.get(3)
	fmt.Println("%v\n", res)
}
