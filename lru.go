package main

type NodeList struct {
	pre, next *NodeList
	k, v      int
}
type LRUCache struct {
	head, tail *NodeList
	m          map[int]*NodeList
	capacity   int
	count      int
}

// head->next ----> tail
// head<---- pre<-tail
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		m:        make(map[int]*NodeList, 0),
	}
	l.head = &NodeList{}
	l.tail = &NodeList{}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.RemoveNode(v)
		this.AddNode(v)
		return v.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.v = value
		this.m[key] = v
		this.RemoveNode(v)
		this.AddNode(v)
	} else {
		if this.count >= this.capacity {
			node := this.tail.pre
			this.RemoveNode(node)
			delete(this.m, node.k)
			this.count--
		}
		node := &NodeList{k: key, v: value}
		this.count++
		this.AddNode(node)
		this.m[key] = node
	}
}
func (this *LRUCache) AddNode(node *NodeList) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}
func (this *LRUCache) RemoveNode(node *NodeList) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
