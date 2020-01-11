package questionsix

// CycleFinder ...
type CycleFinder interface {
	CycleFind() bool
}

// NodeSetter ...
type NodeSetter interface {
	SetNode(node listNoder)
}

type listNoder interface {
	CycleFinder
	NodeSetter
}

// listNode is structure of singly linked list
type listNode struct {
	data int
	next listNoder
}

// HasCycle checks the loop in a linked list
func (l *listNode) CycleFind() bool {
	Slow := l
	Fast := l
	for {
		if Fast.next == nil || Fast.next.(*listNode).next.(*listNode) == nil {
			return false
		}
		Slow = Slow.next.(*listNode)
		Fast = Fast.next.(*listNode).next.(*listNode)

		if Slow == Fast {
			return true
		}
	}
}

// SetNode needs for cycle created (lol)
func (l *listNode) SetNode(node listNoder) {
	l.next = node
}

// NewListNode ...
func NewListNode(data int, node listNoder) listNoder {
	return &listNode{data, node}
}
