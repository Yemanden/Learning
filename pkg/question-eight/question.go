package question_eight

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int
	if root == nil {
		return sum
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}
	if root.Val > L {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R {
		sum += rangeSumBST(root.Right, L, R)
	}
	return sum
}
