package question_eight

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func rangeSumBST(root *treeNode, L int, R int) int {
	var sum int
	if root == nil {
		return sum
	}

	if root.val >= L && root.val <= R {
		sum += root.val
	}
	if root.val > L {
		sum += rangeSumBST(root.left, L, R)
	}
	if root.val <= R {
		sum += rangeSumBST(root.right, L, R)
	}
	return sum
}
