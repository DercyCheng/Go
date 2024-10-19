package Stack_Queue

// isValid 检查给定的括号字符串是否有效
// 有效的括号字符串需满足：
// 1. 左括号必须用相同类型的右括号闭合。
// 2. 左括号必须以正确的顺序闭合。
// 参数:
// - s: 要检查的括号字符串
// 返回值:
// - bool: 如果括号字符串有效，返回 true；否则返回 false。
func isValid(s string) bool {
	stack := make([]rune, 0) // 初始化一个空栈，用于存储左括号
	m := map[rune]rune{}     // 创建一个映射，用于匹配右括号和左括号
	m[']'] = '['
	m['}'] = '{'
	m[')'] = '('
	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c) // 如果是左括号，压入栈中
		} else {
			if len(stack) == 0 {
				return false // 如果栈为空且遇到右括号，返回 false
			}
			peek := stack[len(stack)-1] // 获取栈顶元素
			if peek != m[c] {
				return false // 如果栈顶元素与当前右括号不匹配，返回 false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		}
	}
	return len(stack) == 0 // 如果栈为空，返回 true；否则返回 false
}
