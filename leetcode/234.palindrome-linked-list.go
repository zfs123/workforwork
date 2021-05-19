/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	p, q := head, head.Next
	for q != nil {
		if q.Next != nil {
			q = q.Next.Next
		} else {
			break
		}
		p = p.Next
	}
	r := p.Next
	p.Next = nil
	rhead := reverse(r)
	p, q = head, rhead
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for q != nil {
		n := q.Next
		q.Next = p
		if p == head {
			p.Next = nil
		}
		p = q
		q = n
	}
	return p
}

// @lc code=end

