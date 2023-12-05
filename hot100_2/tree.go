package hot100_2

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
}

// 543. 二叉树的直径 需要先理解二叉树的最大深度
// 输入：root = [1,2,3,4,5]
// 输出：3
// 解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
func diameterOfBinaryTree(root *TreeNode) int {

}

// 101. 对称二叉树
// 给你一个二叉树的根节点 root ， 检查它是否轴对称。
func isSymmetric(root *TreeNode) bool {

}

// 101. 对称二叉树 递归
func IssSymmetric(root *TreeNode) bool {

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
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// 98. 验证二叉搜索树  中序便利，判断是否单调递增
// 有效 二叉搜索树定义如下：
// - 节点的左子树只包含 小于 当前节点的数。
// - 节点的右子树只包含 大于 当前节点的数。
// - 所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {

}

// 230. 二叉搜索树中第K小的元素 中序列遍历，k--
func kthSmallest(root *TreeNode, k int) int {

}

// 230 递归
func kthSmallestRecursion(root *TreeNode, k int) int {

}

// 199. 二叉树的右视图 层序遍历取最右, 最右边输出
// 输入: [1,2,3,null,5,null,4]
// 输出: [1,3,4]
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

}

// 114. 二叉树展开为链表
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
func flatten(root *TreeNode) {

}

// 114 二叉树展开为链表 递归
func flattenRecursion(root *TreeNode) {

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

}

// 124. 二叉树中的最大路径和
// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。
// 该路径 至少包含一个 节点，且不一定经过根节点。
func maxPathSum(root *TreeNode) int {

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

// 前序列遍历
func preorderTraversal(root *TreeNode) []int {
	var list []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		node = node.Right
		stack = stack[:len(stack)-1]
	}
	return list
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
