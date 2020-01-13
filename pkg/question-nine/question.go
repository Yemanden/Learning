package questionnine

// ListReverser ...
type ListReverser interface {
	Reverse() ListReverser
	Getter
	Setter
}

// Getter ...
type Getter interface {
	GetData() []int
}

// Setter ...
type Setter interface {
	SetValue(int)
	SetNextNode(ListReverser)
}

type listNode struct {
	val  int
	next ListReverser
}

// Reverse ...
func (l *listNode) Reverse() ListReverser {
	if l == nil || l.next == nil {
		return l
	}
	var prev ListReverser = nil
	cur := l
	nxt := l.next.(*listNode)
	for nxt != nil {
		cur.next = prev
		prev = cur
		cur = nxt
		if cur.next == nil {
			break
		}
		nxt = cur.next.(*listNode)
	}
	cur.next = prev
	return cur
}

// GetData ...
func (l *listNode) GetData() []int {
	result := []int{l.val}
	if l.next != nil {
		result = append(result, l.next.GetData()...)
	}
	return result
}

// SetValue ...
func (l *listNode) SetValue(value int) {
	l.val = value
}

// SetNextNode ...
func (l *listNode) SetNextNode(next ListReverser) {
	l.next = next
}

// NewListNode ...
func NewListNode() ListReverser {
	return &listNode{}
}
