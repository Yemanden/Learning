package questionsix

type listNoder interface {
	CycleFinder
	NodeSetter
	getter
}

type getter interface {
	getNext() listNoder
}

// CycleFinder ...
type CycleFinder interface {
	CycleFind() bool
}

// NodeSetter ...
type NodeSetter interface {
	SetNode(node listNoder)
	SetValue(int)
}

// listNode is structure of singly linked list
type listNode struct {
	data int
	next listNoder
}

// HasCycle checks the loop in a linked list
func (l *listNode) CycleFind() bool {
	Slow := listNoder(l)
	Fast := listNoder(l)
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
func (l *listNode) SetNode(node listNoder) {
	l.next = node
}

// SetValue ...
func (l *listNode) SetValue(val int) {
	l.data = val
}

func (l *listNode) getNext() listNoder {
	return l.next
}

// NewListNode ...
func NewListNode() listNoder {
	return &listNode{}
}
