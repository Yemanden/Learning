package questioneight

type getter interface {
	getLeftNode() TreeNoder
	getRightNode() TreeNoder
}

// TreeNoder ...
type TreeNoder interface {
	Summator
	Setter
	getter
}

// Setter ...
type Setter interface {
	SetValue(int)
	SetLeftNode(TreeNoder)
	SetRightNode(TreeNoder)
}

// Summator ...
type Summator interface {
	RangeSumma(int, int) int
}

type treeNode struct {
	val   int
	left  TreeNoder
	right TreeNoder
}

// RangeSumma ...
func (t *treeNode) RangeSumma(from int, to int) int {
	var sum int
	if t == nil {
		return sum
	}

	if t.val >= from && t.val <= to {
		sum += t.val
	}
	if t.val > from && t.getLeftNode() != nil {
		sum += t.getLeftNode().RangeSumma(from, to)
	}
	if t.val <= to && t.getRightNode() != nil {
		sum += t.getRightNode().RangeSumma(from, to)
	}
	return sum
}

// SetValue ...
func (t *treeNode) SetValue(value int) {
	t.val = value
}

// SetLeftNode ...
func (t *treeNode) SetLeftNode(node TreeNoder) {
	t.left = node
}

// SetRightNode ...
func (t *treeNode) SetRightNode(node TreeNoder) {
	t.right = node
}

func (t *treeNode) getLeftNode() TreeNoder {
	return t.left
}

func (t *treeNode) getRightNode() TreeNoder {
	return t.right
}

func NewTreeNode() TreeNoder {
	return &treeNode{}
}
