package leetcode_hot100

import "math"

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

// 108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 98. 验证二叉搜索树  中序便利，判断是否单调递增
func isValidBST(root *TreeNode) bool {
	// 第一次第一个节点比较
	pre := math.MinInt
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if node.Val <= pre {
			return false
		}
		// 记录上一个节点
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}

// 230. 二叉搜索树中第K小的元素 中序列遍历，k--
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return 0
}

// 230 递归
func kthSmallestRecursion(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	var res int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if k == 0 {
			return
		}
		k--
		if k == 0 {
			res = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}

// 199. 二叉树的右视图 层序遍历取最右
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, queue[size-1].Val)
		queue = queue[size:]
	}
	return res
}

// 199 递归
func rightSideViewRecursion(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	var res []int
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		depth++
		dfs(node.Right, depth)
		dfs(node.Left, depth)
	}
	dfs(root, 0)
	return res
}

// 114. 二叉树展开为链表
func flatten(root *TreeNode) {
	var res []*TreeNode
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(res); i++ {
		pre := res[i-1]
		pre.Left = nil
		pre.Right = res[i]
	}
}

// 114 二叉树展开为链表 递归
func flattenRecursion(root *TreeNode) {
	var dfs func(*TreeNode)
	var pre *TreeNode
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		dfs(node.Left)
		node.Right = pre
		node.Left = nil
		pre = node
	}
	dfs(root)
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTreeIP(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 后序遍历的最后一个元素是树的根节点
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// 在中序遍历序列中找到根节点的位置
	var rootIndex int
	for i, val := range inorder {
		if val == postorder[len(postorder)-1] {
			rootIndex = i
			break
		}
	}

	// 递归构建左子树和右子树
	root.Left = buildTreeIP(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTreeIP(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

	return root
}

// 437. 路径总和 III 因为任何节点，都可以做为开始的节点
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	countFromRoot := findPath(root, targetSum)
	// 以左子树为根节点继续递归查找
	countFromLeft := pathSum(root.Left, targetSum)

	// 以右子树为根节点继续递归查找
	countFromRight := pathSum(root.Right, targetSum)
	// 返回满足条件的路径总数
	return countFromRoot + countFromLeft + countFromRight
}

func findPath(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val == targetSum {
		count++
	}
	count += findPath(node.Left, targetSum-node.Val)
	count += findPath(node.Right, targetSum-node.Val)
	return count
}

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res = append(res, root.Val)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			tmp := res[i]
			if node.Left == nil && node.Right == nil {
				if tmp == targetSum {
					return true
				}
				continue
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
				res = append(res, node.Left.Val+tmp)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				res = append(res, node.Right.Val+tmp)
			}
		}
		queue = queue[size:]
		res = res[size:]
	}
	return false
}

// 236. 二叉树的最近公共祖先 体会后续遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r == nil {
		return l
	} else if l == nil && r != nil {
		return r
	} else if l != nil && r != nil {
		return root
	} else {
		return nil
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
func inorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
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
