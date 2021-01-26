package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		desc	string
		nums    []int
		target   int
		expect  []int
	}{
		{
			desc: "case1",
			nums: []int{1,2,3,4,5},
			target: 9,
			expect: []int{3, 4},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r := twoSum(tC.nums, tC.target)
			for i, e := range tC.expect {
				if r[i] != e {
					t.Errorf("TesttwoSum failed")
				}
			}
		})
	}
}

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		desc string
		nums []int
		pos int
		expect bool
	}{
		{
			desc: "case1",
			nums: []int{3,2,0,-4},
			pos: 1,
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var nodes []*ListNode
			for i, v := range tt.nums {
				node := new(ListNode)
				node.Val = v
				if i >= 1 {
					nodes[i-1].Next = node
				}
				nodes = append(nodes, node)
			}
			if tt.pos != -1 {
			    nodes[len(nodes) - 1].Next = nodes[tt.pos]
			}
			res := hasCycle(nodes[0])
			if res != tt.expect {
				t.Errorf("TestLinkedListCycle failed")
			}
		})
	}
}