package cryptLogic

import (
	"strings"
)

func Decod(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= '2' && s[i] <= '9' {
			count := int(s[i] - '0')
			if s[i+1] == '(' {
				j := i + 2
				nested := 1
				for ; j < len(s); j++ {
					if s[j] == '(' {
						nested++
					} else if s[j] == ')' {
						nested--
						if nested == 0 {
							break
						}
					}
				}
				substr := s[i+2 : j]
				for k := 0; k < count; k++ {
					result.WriteString(Decod(substr))
				}
				i = j
			} else {
				for k := 0; k < count; k++ {
					result.WriteByte(s[i+1])
				}
				i++
			}
		} else {
			result.WriteByte(s[i])
		}
	}
	return result.String()
}
