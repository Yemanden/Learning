package question_seven

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func mergeTrees(t1 *treeNode, t2 *treeNode) *treeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t3 := &treeNode{}
	t3.val = t1.val + t2.val

	if t1.left != nil {
		if t2.left != nil {
			t3.left = mergeTrees(t1.left, t2.left) // one one
		} else {
			t3.left = mergeTrees(t1.left, nil) // one nil
		}
	} else if t2.left != nil {
		t3.left = mergeTrees(nil, t2.left) // nil one
	}

	if t1.right != nil {
		if t2.right != nil {
			t3.right = mergeTrees(t1.right, t2.right) // one one
		} else {
			t3.right = mergeTrees(t1.right, nil) // one nil
		}
	} else if t2.right != nil {
		t3.right = mergeTrees(nil, t2.right) // nil one
	}

	return t3
}
