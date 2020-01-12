package questionseven

type getter interface {
	getValue() int
	getLeftNode() TreeNoder
	getRightNode() TreeNoder
}

// TreeNoder ...
type TreeNoder interface {
	TreeMerger
	Setter
	getter
}

// Setter ...
type Setter interface {
	SetValue(int)
	SetLeftNode(TreeNoder)
	SetRightNode(TreeNoder)
}

// TreeMerger ...
type TreeMerger interface {
	TreeMerge(TreeNoder) TreeNoder
}

type treeNode struct {
	val   int
	left  TreeNoder
	right TreeNoder
}

// TreeMerge ...
func (t *treeNode) TreeMerge(otherTree TreeNoder) TreeNoder {
	if t == nil {
		return otherTree
	}
	if otherTree == nil {
		return t
	}

	t3 := &treeNode{}
	t3.val = t.val + otherTree.getValue()

	if t.left != nil {
		if otherTree.getLeftNode() != nil {
			t3.left = t.left.TreeMerge(otherTree.getLeftNode()) // one one
		} else {
			t3.left = t.left // one nil
		}
	} else {
		t3.left = otherTree.getLeftNode() // one nil
	}

	if t.right != nil {
		if otherTree.getRightNode() != nil {
			t3.right = t.right.TreeMerge(otherTree.getRightNode()) // one one
		} else {
			t3.right = t.right // one nil
		}
	} else {
		t3.right = otherTree.getRightNode() // nil one
	}

	return t3
}

// SetValue ...
func (t *treeNode) SetValue(val int) {
	t.val = val
}

// SetLeftNode ...
func (t *treeNode) SetLeftNode(node TreeNoder) {
	t.left = node
}

// SetRightNode ...
func (t *treeNode) SetRightNode(node TreeNoder) {
	t.right = node
}

func (t *treeNode) getValue() int {
	return t.val
}

func (t *treeNode) getLeftNode() TreeNoder {
	return t.left
}

func (t *treeNode) getRightNode() TreeNoder {
	return t.right
}

// NewTreeNode ...
func NewTreeNode() TreeNoder {
	return &treeNode{}
}
