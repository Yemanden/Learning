package question_eight

// TreeSummarizer ...
type TreeSummarizer interface {
	TreeSummarize(int, int) int
}

// TreeInitializationer ...
type TreeInitializationer interface {
	TreeInitialization(int, treeNoder, treeNoder)
}

type treeNoder interface {
	TreeSummarizer
	TreeInitializationer
}

type treeNode struct {
	val   int
	left  treeNoder
	right treeNoder
}

// TreeSummarize ...
func (t *treeNode) TreeSummarize(L int, R int) int {
	var sum int
	if t == nil {
		return sum
	}

	if t.val >= L && t.val <= R {
		sum += t.val
	}
	if t.val > L {
		sum += t.left.TreeSummarize(L, R)
	}
	if t.val <= R {
		sum += t.right.TreeSummarize(L, R)
	}
	return sum
}

// TreeInitialization ...
func (t *treeNode) TreeInitialization(val int, left, right treeNoder) {
	t.val = val
	t.left = left
}

/// NewTreeNode ...
func NewTreeNode() treeNoder {
	return &treeNode{}
}
