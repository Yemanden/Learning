package questioneight

// TreeNoder ...
type TreeNoder interface {
	Summator
	Setter
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
	if t.val > from && t.left != nil {
		sum += t.left.RangeSumma(from, to)
	}
	if t.val <= to && t.right != nil {
		sum += t.right.RangeSumma(from, to)
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

func NewTreeNode() TreeNoder {
	return &treeNode{}
}
