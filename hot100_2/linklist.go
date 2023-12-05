package hot100_2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。
// 如果两个链表不存在相交节点，返回 null 。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA := headA
	hB := headB
	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}
		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return hA
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 206. 反转链表 递归
func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return cur
}

func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Println(slow.Val)
	cur := reverseListNode(slow)

	for cur != nil {
		fmt.Println(head.Val, cur.Val)
		if head.Val != cur.Val {
			return false
		}
		head = head.Next
		cur = cur.Next
	}

	return true
}

func isPalindromeV2(head *ListNode) bool {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i, j := 0, len(values)-1; i < j; {
		if values[i] != values[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
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
	isCycle := func(h *ListNode) (bool, *ListNode) {
		slow, fast := h, h
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true, slow
			}
		}
		return false, nil
	}
	valid, slow := isCycle(head)
	if valid {
		for head != slow {
			head = head.Next
			slow = slow.Next
		}
		return slow
	}
	return nil
}

// 21. 合并两个有序链表
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
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
	dummy := &ListNode{}
	plus := 0
	cur := dummy
	for l1 != nil && l2 != nil {
		num := l1.Val + l2.Val + plus
		node := &ListNode{Val: (num) % 10}
		plus = num / 10
		cur.Next = node
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		num := l1.Val + plus
		node := &ListNode{Val: (num) % 10}
		plus = num / 10
		cur.Next = node
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		num := l2.Val + plus
		node := &ListNode{Val: (num) % 10}
		plus = num / 10
		cur.Next = node
		cur = cur.Next
		l2 = l2.Next
	}
	if plus != 0 {
		node := &ListNode{Val: plus % 10}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, cur := dummy, dummy
	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	for cur != nil && cur.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 24. 两两交换链表中的节点
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for cur := head; cur != nil && cur.Next != nil; {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		q := next.Next
		next.Next = cur
		pre = cur
		cur = q
	}
	return dummy.Next
}

// 24. 两两交换链表中的节点 递归
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairsRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	tmp := swapPairsRecursion(next.Next)
	head.Next = tmp
	next.Next = head
	return next
}

// 25. K 个一组翻转链表
// 输入：head = [1,2,3,4,5], k = 2
// 输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	preHead := dummy
	for head != nil {
		oHead := head
		groupLastTail := head
		for i := 1; i < k && groupLastTail != nil; i++ {
			groupLastTail = groupLastTail.Next
		}
		if groupLastTail == nil {
			break
		}
		groupHead := groupLastTail.Next
		groupLastTail.Next = nil
		revertHead := reverseList(head)
		preHead.Next = revertHead
		oHead.Next = groupHead
		head = groupHead
		preHead = oHead
	}
	return dummy.Next
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
	m := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = &Node{Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Next != nil {
			m[cur].Next = m[cur.Next]
		}
		if cur.Random != nil {
			m[cur].Random = m[cur.Random]
		}
	}
	return m[head]
}

// 148. 排序链表
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
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
	var res *ListNode
	for i := range lists {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}
