package Stack_Queue

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		atoi, err := strconv.Atoi(token)
		if err != nil {
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch atoi {
			case '+':
				stack = append(stack, num1+num2)
			case '-':
				stack = append(stack, num1-num2)
			case '*':
				stack = append(stack, num1*num2)
			case '/':
				stack = append(stack, num1/num2)
			}
		} else {
			stack = append(stack, atoi)
		}
	}
	return stack[0]
}
