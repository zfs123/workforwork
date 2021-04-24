/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	return mergeSortList(head)
}

func mergeSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	rhead := mid.Next
	mid.Next = nil
	left := mergeSortList(head)
	right := mergeSortList(rhead)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	n := dummy
	p, q := l1, l2
	for p != nil && q != nil {
		if p.Val <= q.Val {
			n.Next = p
			n = n.Next
			p = p.Next
		} else {
			n.Next = q
			n = n.Next
			q = q.Next
		}
	}
	if p == nil {
		n.Next = q
	} else {
		n.Next = p
	}
	return dummy.Next
}

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