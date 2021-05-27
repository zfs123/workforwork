/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	top := -1
	for _, it := range tokens {
		var a, b int
		switch it {
		case "+", "-", "*", "/":
			a = stack[top]
			top--
			b = stack[top]
			stack = stack[:top]
			top--
		default:
			v, _ := strconv.Atoi(it)
			stack = append(stack, v)
			top++
			continue
		}
		var tmp int
		switch it {
		case "+":
			tmp = a + b
		case "-":
			tmp = b - a
		case "*":
			tmp = a * b
		case "/":
			tmp = b / a
		}
		stack = append(stack, tmp)
		top++
	}
	return stack[0]
}

// @lc code=end

