/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{Val: 0}
	p, q, r := l1, l2, dummy
	for p != nil && q != nil {
		if p.Val <= q.Val {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}
	if p == nil {
		r.Next = q
	} else {
		r.Next = p
	}
	return dummy.Next
}