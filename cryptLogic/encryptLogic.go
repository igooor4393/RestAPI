package cryptLogic

import (
	"fmt"
	"strings"
)

func Encode(s string) string {
	var result strings.Builder
	count := 1
	insideSeq := false
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			count++
		} else {
			if count > 1 {
				result.WriteString(fmt.Sprintf("%d", count))
				if !insideSeq {
					result.WriteRune('(')
					insideSeq = true
				}
			} else {
				result.WriteString(fmt.Sprintf("%d", count))
			}
			result.WriteByte(s[i])
			if i < len(s)-2 && (s[i] != s[i+2] || i == len(s)-2) && insideSeq {
				result.WriteRune(')')
				insideSeq = false
			}
			count = 1
		}
	}
	return result.String()
}
