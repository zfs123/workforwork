/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p, q := head, head.Next
	for q != nil {
		n := q.Next
		if p == head {
			p.Next = nil
		}
		q.Next = p
		p, q = q, n
	}
	return p
}