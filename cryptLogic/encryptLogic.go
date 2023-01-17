package cryptLogic

import (
	"fmt"
)

func Encode(s string) string {
	var count int
	var result []byte
	for i := 0; i < len(s); i++ {
		count++
		if i+1 >= len(s) || s[i] != s[i+1] {
			if count > 1 {
				result = append(result, []byte(fmt.Sprintf("%d", count))...)
			}
			result = append(result, s[i])
			count = 0
		}
	}
	var stack []byte
	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			stack = append(stack, result[i])
		} else {
			if len(stack) > 0 {
				var num int
				for _, x := range result[i:] {
					if x >= '0' && x <= '9' {
						num = num*10 + int(x-'0')
					} else {
						break
					}
				}
				stack = append(stack, '(')
				for j := 0; j < num; j++ {
					for _, x := range stack[len(stack)-len(fmt.Sprintf("%d", num)):] {
						stack = append(stack, x)
					}
				}
				stack = stack[:len(stack)-len(fmt.Sprintf("%d", num))]
				stack = append(stack, ')')
				i += len(fmt.Sprintf("%d", num))
			}
		}
	}

	return string(result)
}
