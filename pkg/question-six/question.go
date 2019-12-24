package question_six

// listNode is structure of singly linked list
type listNode struct {
	Data int
	next *listNode
}

// HasCycle checks the loop in a linked list
func HasCycle(h *listNode) bool {
	Slow := h
	Fast := h
	for {
		if Fast.next == nil || Fast.next.next == nil {
			return false
		}
		Slow = Slow.next
		Fast = Fast.next.next

		if Slow == Fast {
			return true
		}
	}

	return false
}
