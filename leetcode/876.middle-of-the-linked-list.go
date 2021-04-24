/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head // return the back one, [1,2,3,4] -> 3
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	return p
}

// return the front one
func findMid(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head.Next
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	return p
}