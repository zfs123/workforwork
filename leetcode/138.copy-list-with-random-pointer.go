/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p, q := head, head
	var nh *Node
	for q != nil {
		r := &Node{Val: q.Val}
		if q != head {
			p.Next = r
		} else {
			nh = r
		}
		q = q.Next
		p = r
	}

	p, np := head, nh
	for p != nil {
		t, nt := head, nh
		for t != nil {
			if t == p.Random {
				break
			}
			t = t.Next
			nt = nt.Next
		}
		np.Random = nt
		p = p.Next
		np = np.Next
	}

	return nh
}