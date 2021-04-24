/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return
    }a
    mid := findMid(head)
    t := mid.Next
    mid.Next = nil
    rhead := reverseList(t)
    p, q, last := head, rhead, rhead
    
    for p != nil && q != nil {
        n := p.Next
        p.Next = q
        rn := q.Next
        q.Next = n
        last = q
        p = n
        q = rn
    }
    if p == nil {
        last.Next = q
    }
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
		p = q
        q = n
	}
	return p
}