package hot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(root *TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)

	}
	traversal(root)

	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	sum := 0
	for len(queue) > 0 {
		size := len(queue)
		sum++
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}

	return sum
}
