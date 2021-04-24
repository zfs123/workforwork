/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, nil}
	dummy.Next = head
	head = dummy

	front, p, q := head, head, head.Next
	count := 1
	for q != nil && count <= left {
		front = p
		p = p.Next
		q = q.Next
		count++
	}
	back := p
	for q != nil && count <= right {
		n := q.Next
		q.Next = p
		p = q
		q = n
		count++
	}
	front.Next = p
	back.Next = q
	return dummy.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	p, q := head, head.Next
	count := 2
	for q != nil && count < left {
		p = p.Next
		q = q.Next
		count = count + 1
	}

	front := &ListNode{0, nil}
	var back *ListNode
	if left == 1 {
		front.Next = head
		back = head
		count = count - 1
	} else {
		front = p
		back = q
		p = q
		q = q.Next
	}

	if q == nil {
		return nil
	}

	for q != nil && count < right {
		n := q.Next
		q.Next = p
		p = q
		q = n
		count = count + 1
	}

	front.Next = p
	back.Next = q
	if left == 1 {
		return front.Next
	} else {
		return head
	}
}