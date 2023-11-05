package leetcode_hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

// 160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hA, hB := headA, headB
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
		q := cur.Next
		cur.Next = pre
		pre = cur
		cur = q
	}
	return pre
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-i-1] {
			return false
		}
	}
	return true
}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func isCycle(slow, fast *ListNode) (bool, *ListNode) {
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	isC, node := isCycle(slow, fast)
	if isC {
		for head != nil {
			if head == node {
				return node
			}
			head = head.Next
			node = node.Next
		}
	}
	return nil
}

// 21. 合并两个有序链表
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	plus := 0
	for l1 != nil || l2 != nil {
		sum := plus
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		val := sum % 10
		plus = sum / 10
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if plus > 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return dummy.Next
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	// size := 0
	// 这里还是少走了一位！！！
	//for fast != nil {
	//	fast = fast.Next
	//	size++
	//	if size == n {
	//		break
	//	}
	//}
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil && cur.Next.Next != nil; {
		q := cur.Next
		p := cur.Next.Next
		cur.Next = p
		cur = q
		q.Next = p.Next
		p.Next = q
	}
	return dummy.Next
}

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	prevGroupTail := dummy
	for head != nil {
		groupHead := head
		groupTail := head
		for i := 0; i < k && groupTail != nil; i++ {
			groupTail = groupTail.Next
		}
		if groupTail == nil {
			break
		}
		nextGroupHead := groupTail.Next
		groupTail.Next = nil

		prevGroupTail.Next = reverseList(groupHead)
		groupHead.Next = nextGroupHead
		prevGroupTail = groupHead
		head = nextGroupHead
	}
	return dummy.Next
}
