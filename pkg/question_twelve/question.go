package questiontwelve

type listNode struct {
	val  int
	next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	l3 := &listNode{0, nil}

	nxt1 := l1
	nxt2 := l2
	nxt3 := l3
	carry := 0

	for nxt1 != nil || nxt2 != nil {
		var sum = 0
		if nxt1 != nil {
			sum += nxt1.val
			nxt1 = nxt1.next
		}
		if nxt2 != nil {
			sum += nxt2.val
			nxt2 = nxt2.next
		}
		sum += carry
		carry = 0
		if sum >= 10 {
			carry = sum / 10
			sum %= 10
		}
		nxt3.val = sum
		if nxt1 != nil || nxt2 != nil || carry != 0 {
			nxt3.next = &listNode{carry, nil}
			nxt3 = nxt3.next
		}
	}
	return l3
}
