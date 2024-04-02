package 总结归类

import "math"

type Node struct {
	k, v      int
	pre, next *Node
	Children  []*Node
}
type LRUCache struct {
	head, tail *Node
	m          map[int]*Node
	capacity   int
	size       int
}

func ConstructorV1(capacity int) LRUCache {
	l := LRUCache{}
	l.capacity = capacity
	l.m = make(map[int]*Node)
	l.head = &Node{}
	l.tail = &Node{}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if n, ok := l.m[key]; ok {
		l.DelNode(n)
		l.AddNode(n)
		return n.v
	} else {
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.m[key]; ok {
		l.DelNode(n)
		n.v = value
		l.AddNode(n)
		l.m[key] = n
	} else {
		if l.size == l.capacity {
			tmp := l.tail.pre
			l.DelNode(tmp)
			delete(l.m, tmp.k)
			l.size--
		}
		n := &Node{k: key, v: value}
		l.AddNode(n)
		l.m[key] = n
		l.size++
	}
}

func (l *LRUCache) AddNode(n *Node) {
	t := l.head.next
	l.head.next = n
	n.pre = l.head
	n.next = t
	t.pre = n
}
func (l *LRUCache) DelNode(n *Node) {
	n.pre.next = n.next
	n.next.pre = n.pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// k个一组翻转链表
//
//		1 2 3 4
//		2 1 4 3
//	    1      2       3
//
// cur ori    node    nextHead
func reverseNode(node *ListNode) *ListNode {
	var pre *ListNode
	for node != nil {
		q := node.Next
		node.Next = pre
		pre = node
		node = q
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := &ListNode{}
	pre := cur
	for head != nil {
		originHead := head
		for i := 1; i < k && head != nil; i++ {
			head = head.Next
		}
		if head == nil {
			return pre.Next
		}
		nextHead := head.Next
		head.Next = nil
		node := reverseNode(originHead)
		cur.Next = node
		head = nextHead
		originHead.Next = nextHead
		cur = originHead
	}
	return pre.Next
}

// 接雨水
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
func trap(height []int) int {
	res := 0
	stack := make([]int, len(height)+1)
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			l := stack[len(stack)-1]
			w := i - l - 1
			h := min(height[l], height[i]) - height[index]
			res += w * h

		}
		stack = append(stack, i)
	}
	return res
}
func trapV2(height []int) int {
	n := len(height)
	res := 0
	l, r := make([]int, len(height)), make([]int, len(height))
	l[0] = height[0]
	r[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		l[i] = max(height[i], l[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		r[i] = max(height[i], r[i+1])
	}
	for i := 0; i < n; i++ {
		res += min(l[i], r[i]) - height[i]
	}
	return res
}

// 84. 柱状图中最大的矩形
func largestRectangleArea(heights []int) int {
	res := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 前一个做为起始点
			l := stack[len(stack)-1]
			w := i - l - 1
			res = max(res, h*w)
		}
		stack = append(stack, i)
	}
	return res
}

// 二维矩阵搜索一个数
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		l := 0
		r := n - 1
		for l <= r {
			mid := (r-l)/2 + l
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

// 卖股票多次机会
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
func maxProfit(prices []int) int {
	// - dp[i][0]表示持有这只股票的最大现金
	// - dp[i][1]表示不持有这只股票的最大现金
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树输出最小路径
// LCR 051. 二叉树中的最大路径和
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		res = max(res, node.Val+l+r)
		if l > r {
			return max(l+node.Val, 0)
		}
		return max(r+node.Val, 0)
	}
	dfs(root)
	return res
}

// 194. 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == q.Val || root.Val == p.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return root
}

// 三数之和
// 最小栈实现
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		top := this.minStack[len(this.minStack)-1]
		this.minStack = append(this.minStack, min(top, x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		// 返回主栈的栈顶元素
		return this.stack[len(this.stack)-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if len(this.stack) > 0 {
		return this.minStack[len(this.minStack)-1]
	}
	return -1
}

// 两个链表有没有交点

// 多叉树的最大深度
func maxDepth(root *Node) int {
	res := 0
	var dfs func(root *Node, deep int) int
	dfs = func(root *Node, deep int) int {
		if root == nil {
			return deep
		}
		res = max(res, deep)
		for _, v := range root.Children {
			dfs(v, deep+1)
		}
		return deep
	}
	dfs(root, 1)
	return res
}
