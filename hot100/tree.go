package leetcode_hot100

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

// 后序遍历，先序遍历反转
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// InorderTraversal 中序遍历 左中右 没有过，明天继续写
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return res
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		size := len(queue)
		arr := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[i]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		res = append(res, arr)
	}
	return res
}
