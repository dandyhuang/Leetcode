package link_list

// 203.移除链表元素
// 题意：删除链表中等于给定值 val 的所有节点。
// 示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]
// 示例 2： 输入：head = [], val = 1 输出：[]
// 示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
type ListNode struct {
	Val  int
	Next *ListNode
}

// 亚节点
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 删除头节点情况
func removeElementsV2(head *ListNode, val int) *ListNode {
	// 头节点情况
	for head != nil && head.Val == val {
		head = head.Next
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 707.设计链表
type MyLinkedList struct {
	dummy *Node // 虚拟节点
	size  int
}
type Node struct {
	val  int
	next *Node
}

func Constructor() MyLinkedList {
	node := &Node{
		-1,
		nil,
	}
	return MyLinkedList{node, 0}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.dummy.next
	for i := 0; i < this.size; i++ {
		if i == index {
			return cur.val
		}
		cur = cur.next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	node := &Node{val: val}
	cur := this.dummy
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	node.next = cur.next
	cur.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.dummy
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// 输入
//["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
//[[], [1], [3], [1, 2], [1], [1], [1]]
//输出
//[null, null, null, null, 2, null, 3]

// 206.反转链表
// 题意：反转一个单链表。
// 示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL
// pre head next
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// 从后向前递归
func recursionReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// cur = 5, head = 4
	cur := recursionReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for head != nil && head.Next != nil {
		next := head.Next
		cur.Next = next
		head.Next = next.Next
		cur = head
		next.Next = head
		head = head.Next
	}
	return dummy.Next
}
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsRecursion(next.Next)
	next.Next = head
	return next
}

// 25. K 个一组翻转链表
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {

}

// 19.删除链表的倒数第N个节点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
// 输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5]
// 示例 2：
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	count := 0
	cur := dummy
	for cur != nil {
		cur = cur.Next
		if count == n {
			break
		}
		count++
	}
	pre := dummy
	for cur != nil {
		cur = cur.Next
		pre = pre.Next
	}
	pre.Next = pre.Next.Next
	return dummy.Next
}

// 面试题 02.07. 链表相交
// 同：160.链表相交
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {

}
