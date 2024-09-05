package main

import (
	"fmt"
	"strconv"
	"sync"
)

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
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)

		}()
	}
}
func letterCombinations(digits string) []string {
	//m := map[byte]string{
	//	'1': "",
	//	'2': "abc",
	//	'3': "def",
	//	'4': "ghi",
	//	'5': "jkl",
	//	'6': "mno",
	//	'7': "pqrs",
	//	'8': "tuv",
	//	'9': "wxyz",
	//}
	var res []string
	//var dfs func(start int, str string)
	//dfs = func(start int, str string) {
	//	if len(str) == len(digits) {
	//		res = append(res, str)
	//		return
	//	}
	//	for i := 0; i < len(digits); i++ {
	//		tmp := m[digits[i]][start]
	//		dfs(start+1, str+m[digits[i]][start])
	//	}
	//}
	//dfs(0, "")
	return res
}
func isNormalIp(s string, start, end int) bool {
	if s[0] == '0' && end-start+1 > 1 {
		return false
	}
	num, _ := strconv.Atoi(s[start : end+1])
	if num > 255 {
		return false
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	res := 0
	sum := 0
	for i := range nums {
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		res = max(res, sum)
	}
	return res
}
func maxSubArrayV2(nums []int) int {
	res := 0
	dp := make([]int, len(nums))
	dp[0] = max(nums[0], 0)
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], 0)
		res = max(res, dp[i])
	}
	return res
}
func main() {
	var (
		wg    = &sync.WaitGroup{}
		count int
	)
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func() {
			defer wg.Done()
			count += i
		}()
	}
	wg.Wait()
	fmt.Println(count)

	//a, err := strconv.Atoi("place")
	//if err != nil {
	//	fmt.Println(a, err)
	//}
	//res := maxSubArrayV2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	//fmt.Println("res:", res)
	//	//x := 7
	//	//z:= x / 2.0
	//fmt.Printf("%v, %v, %v, %v\n", x/2.0, float64(x)/float64(2), x, z)
	//s := "13ewer"
	//fmt.Println(s[1:4])
	//DeferClosureLoopV1()
	//// 创建示例链表: 1 -> 2 -> 3 -> 4
	//// head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	//// 创建示例链表: 1 -> 2 -> 3 -> 4 -> 5
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	//
	//// 重新排列链表
	//reorderList(head)
	//
	//// 打印结果
	//printList(head)
}
