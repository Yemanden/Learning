package question_six

// ListNode is structure of singly linked list
type ListNode struct {
	Data int
	Next *ListNode
}

// HasCycle checks the loop in a linked list
func HasCycle(head *ListNode) bool {
	Slow := head
	Fast := head
	for true {
		if Fast.Next == nil || Fast.Next.Next == nil {
			return false
		}
		Slow = Slow.Next
		Fast = Fast.Next.Next

		if Slow == Fast {
			return true
		}
	}

	return false
}
