package tt_LeetcodeTop

type Node struct {
	value, key int
	pre, next  *Node
}

type LRUCache struct {
	size, capacity int
	head           *Node
	tail           *Node
	m              map[int]*Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{}
	l.capacity = capacity
	l.size = 0
	l.head = &Node{}
	l.tail = &Node{}
	l.head.next = l.tail
	l.tail.pre = l.head
	l.m = make(map[int]*Node)
	return l
}

func (l *LRUCache) Get(k int) int {
	if v, ok := l.m[k]; ok {
		l.DeleteNode(v)
		l.AddNode(v)
		return v.value
	}
	return -1
}

func (l *LRUCache) Put(k int, v int) {
	if n, ok := l.m[k]; ok {
		delete(l.m, n.key)
		l.DeleteNode(n)
	} else {
		if l.capacity == l.size {
			node := l.tail.pre
			l.DeleteNode(node)
			delete(l.m, node.key)
			l.size--
		}
		l.size++
	}
	n := &Node{key: k, value: v}
	l.m[k] = n
	l.AddNode(n)
}

func (l *LRUCache) DeleteNode(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

func (l *LRUCache) AddNode(n *Node) {
	n.next = l.head.next
	n.pre = l.head
	// 这里很关键
	l.head.next.pre = n
	l.head.next = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
