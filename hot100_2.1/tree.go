package hot100_2

import (
	"math"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {

}

// 94. 二叉树的前序遍历
func preOrderTraversal(root *TreeNode) []int {
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

// 543. 二叉树的直径 需要先理解二叉树的最大深度
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left) + 1
		r := dfs(node.Right) + 1
		res = max(res, l+r)
		return max(l, r)
	}
	dfs(root)
	// 左边+右边，算的是边
	return res - 2
}

// 226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	root.Left = r
	root.Right = l
	return root
}

// 101. 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

func mirror(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return mirror(l.Left, r.Right) && mirror(l.Right, r.Left)
}

// 101. 对称二叉树 递归
func IssSymmetric(root *TreeNode) bool {

}

// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
// 从底至顶
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		// 增加判断，继续将错误往上抛
		if l == -1 {
			return -1
		}
		r := dfs(node.Right)
		// 增加判断，继续往上抛
		if r == -1 {
			return -1
		}
		// 如果中间有 >1的情况，会导致l或者r返回-1。 后续在递归l-r的时候，计算就会不准确了
		if abs(l-r) > 1 {
			return -1
		}
		return max(l, r) + 1
	}
	return dfs(root) != -1
}

// 从顶至底
func isBalancedV2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		return max(l, r) + 1
	}
	return abs(dfs(root.Left)-dfs(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 108. 将有序数组转换为二叉搜索树
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	l := sortedArrayToBST(nums[:mid])
	r := sortedArrayToBST(nums[mid+1:])
	return &TreeNode{
		Left:  l,
		Right: r,
		Val:   nums[mid],
	}
}

// 98. 验证二叉搜索树  中序便利，判断是否单调递增
// 有效 二叉搜索树定义如下：
// - 节点的左子树只包含 小于 当前节点的数。
// - 节点的右子树只包含 大于 当前节点的数。
// - 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre > node.Val {
			return false
		}
		pre = node.Val
		root = node.Right
	}
	return true
}

// 230. 二叉搜索树中第K小的元素 中序列遍历，k--
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	count := 1
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if count == k {
			return node.Val
		}
		count++
		root = node.Right
	}
	return -1
}

// 230 递归
func kthSmallestRecursion(root *TreeNode, k int) int {

}

// 199. 二叉树的右视图 层序遍历取最右, 最右边输出
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
func rightSideView(root *TreeNode) []int {

}

// 199 递归
func rightSideViewRecursion(root *TreeNode) []int {

}

// 114. 二叉树展开为链表 左中右，先序遍历
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
func flatten(root *TreeNode) {

}

// 114 二叉树展开为链表 递归
func flattenRecursion(root *TreeNode) {
	var pre *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 先从right开始遍历
		flatten(root.Right)
		flatten(root.Left)
		root.Left = nil
		root.Right = pre
		pre = root
	}
	dfs(root)
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {

}

// 106. 从中序与后序遍历序列构造二叉树
func buildTreeIP(inorder []int, postorder []int) *TreeNode {

}

// 437. 路径总和 III 因为任何节点，都可以做为开始的节点
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
func pathSum(root *TreeNode, targetSum int) int {
}

// 112 路径总和
// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，
// 这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
func hasPathSum(root *TreeNode, targetSum int) bool {

}

// 236. 二叉树的最近公共祖先 体会后续遍历
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil && r != nil {
		return r
	}
	if l != nil && r == nil {
		return l
	}
	return nil
}

// 124. 二叉树中的最大路径和
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。
// 该路径 至少包含一个 节点，且不一定经过根节点。
func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		sum := l + r + node.Val
		res = max(res, sum)
		maxSum := node.Val + max(l, r)
		return max(maxSum, 0)
	}
	dfs(root)
	return res
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
