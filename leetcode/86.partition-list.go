/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy, dummy2 := &ListNode{0, nil}, &ListNode{0, nil}
	dummy.Next = head
	head = dummy
	p, q, r := head, head.Next, dummy2
	for q != nil {
		if q.Val < x {
			r.Next = q
			p.Next = q.Next
			q.Next = nil
			q = p.Next
			r = r.Next
		} else {
			p = p.Next
			q = q.Next
		}
	}
	r.Next = dummy.Next
	return dummy2.Next
}