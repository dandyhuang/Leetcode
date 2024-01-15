package hot100_2

type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
// 如果两个链表不存在相交节点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
}
dummy, cur, q

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var	dummy *ListNode
	for cur:= head;cur!= nil; {
		q:=cur.Next
		cur.Next = dummy
		dummy = cur
		cur = q
	}
	return dummy
}

// 206. 反转链表 递归
func reverseListRecursion(head *ListNode) *ListNode {

}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {


}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	fast, slow:= head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			 return true
		}
	}
	return false
}

// 142. 环形链表 II
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
func detectCycle(head *ListNode) *ListNode {
	hasCycle:= false
	fast, slow:= head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
	 return nil
	}

	for slow != head {
		slow = slow.Next
		head = head.Next
	}
	return slow
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur:= dummy
	for list1 != nil  && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next  = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
// 2. 两数相加
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 没哑节点，想下两个节点，删除倒数第二个节点怎么办
	dummy:=&ListNode{Next: head}
	slow, fast:= dummy, dummy
	for i:=0; i < n ;i++ {
		fast = fast.Next
	}
	// 运行一下，就知道fast.Next也需要不等于nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 24. 两两交换链表中的节点
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	// 没哑节点，返回的时候，不知道怎么返回
	cur:= head
	dump:= &ListNode{Next: head}
	pre := dump
	for cur != nil && cur.Next !=nil {
		next:=cur.Next
		cur.Next = next.Next
		pre.Next = next
		pre = cur
		cur = next.Next
		next.Next = pre
	}
	return dump.Next
}

// 24. 两两交换链表中的节点 递归
func swapPairsRecursion(head *ListNode) *ListNode {

}

// 25. K 个一组翻转链表
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 138. 随机链表的复制
// 输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
func copyRandomList(head *Node) *Node {
}

// 148. 排序链表
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
func sortList(head *ListNode) *ListNode {

}

// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 23. 合并 K 个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
}
