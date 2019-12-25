package question_nine

type listNode struct {
	val  int
	next *listNode
}

func reverseList(head *listNode) *listNode {
	if head == nil {
		return head
	}
	var prev *listNode = nil
	cur := head
	nxt := head.next
	for nxt != nil {
		cur.next = prev
		prev = cur
		cur = nxt
		nxt = cur.next
	}
	cur.next = prev
	return cur
}
