/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, i := range nums {
		m[i] = idx
	}
	for idx, i := range nums {
		val, ok := m[target - i]
		if ok && idx != val {
			return []int{idx, val}
		}
	}
	return []int{}
}
// @lc code=end

