package questionsix

type getter interface {
	getNext() ListNoder
}

// ListNoder ...
type ListNoder interface {
	CycleFinder
	NodeSetter
	getter
}

// CycleFinder ...
type CycleFinder interface {
	CycleFind() bool
}

// NodeSetter ...
type NodeSetter interface {
	SetNode(node ListNoder)
	SetValue(int)
}

// listNode is structure of singly linked list
type listNode struct {
	data int
	next ListNoder
}

// CycleFind checks the loop in a linked list
func (l *listNode) CycleFind() bool {
	Slow := ListNoder(l)
	Fast := ListNoder(l)
	for {
		if Fast.getNext() == nil || Fast.getNext().getNext() == nil {
			return false
		}
		Slow = Slow.getNext()
		Fast = Fast.getNext().getNext()

		if Slow == Fast {
			return true
		}
	}
}

// SetNode needs for cycle created (lol)
func (l *listNode) SetNode(node ListNoder) {
	l.next = node
}

// SetValue ...
func (l *listNode) SetValue(val int) {
	l.data = val
}

func (l *listNode) getNext() ListNoder {
	return l.next
}

// NewListNode ...
func NewListNode() ListNoder {
	return &listNode{}
}
