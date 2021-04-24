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

    f, p, q := head, head, head.Next
    for q != nil {
        if p.Val == q.Val {
            for q.Next != nil {
                if q.Next.Val == p.Val {
                    q = q.Next
                } else {
                    break
                }
            }
            if p == head {
                head = q.Next
                f = head
            } else {
                f.Next = q.Next
            }
            p = q.Next
            q.Next = nil
            if p != nil {
                q = p.Next
            } else {
                break
            }
        } else {
            f = p
            p = p.Next
            q = q.Next
        }

    }
    return head
}