/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    p, q:= head, head.Next
    for q != nil {
        if p.Val == q.Val {
            p.Next = q.Next
            q.Next = nil
            q = p.Next
        } else {
            p = p.Next
            q = q.Next
        }
    }
    return head
}