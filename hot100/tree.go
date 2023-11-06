package leetcode_hot100

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

// 543. 二叉树的直径 需要先理解二叉树的最大深度
func depth(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	l := depth(node.Left, ans)
	r := depth(node.Right, ans)
	*ans = max(*ans, l+r)
	// 返回当前节点为根的子树的深度
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	depth(root, &ans)
	return ans
}

// 101. 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root, root}

	for len(queue) > 0 {
		l := queue[0]
		r := queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left, r.Right)
		queue = append(queue, l.Right, r.Left)
	}

	return true
}

// 101. 对称二叉树 递归
func IssSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
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
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
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
