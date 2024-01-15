package main

import "fmt"

func DeferLoopV20() {
	// i变量在每次调用中被保存了下来
	for i := 0; i < 10; i++ {
		defer func(i int) {
			println(i)
		}(i)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) {
	for l1 != nil && l2 != nil {
		l1Next := l1.Next
		l1.Next = l2
		l1 = l1Next

		l2Next := l2.Next
		l2.Next = l1
		l2 = l2Next
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	revert := func(node *ListNode) *ListNode {
		var dummy *ListNode
		for cur := node; cur != nil; {
			next := cur.Next
			cur.Next = dummy
			dummy = cur
			cur = next
		}
		return dummy
	}
	next := revert(slow)
	// mergeTwoLists(head, next)
	for head != nil && next != nil {
		l1Next := head.Next
		l2Next := next.Next

		head.Next = next
		head = l1Next

		next.Next = head
		next = l2Next
	}
}

// printList 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
func main() {
	// 创建示例链表: 1 -> 2 -> 3 -> 4
	// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	// 创建示例链表: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	// 重新排列链表
	reorderList(head)

	// 打印结果
	printList(head)
}
