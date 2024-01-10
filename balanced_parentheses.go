package sprint

func BalancedParentheses(input string) bool {
	var stack []rune

	for _, ch := range input {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (ch == ')' && last != '(') || (ch == ']' && last != '[') || (ch == '}' && last != '{') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
