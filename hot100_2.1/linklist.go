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

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {

}

// 206. 反转链表 递归
func reverseListRecursion(head *ListNode) *ListNode {

}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {

}

// 141. 环形链表
func hasCycle(head *ListNode) bool {
}

// 142. 环形链表 II
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。
func detectCycle(head *ListNode) *ListNode {

	return nil
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

}

// 24. 两两交换链表中的节点
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
func swapPairs(head *ListNode) *ListNode {

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
