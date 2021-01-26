package leetcode


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

// best solution
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}