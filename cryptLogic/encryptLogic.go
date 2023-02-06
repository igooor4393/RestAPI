package cryptLogic

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Encode(s string) string {

	var result string
	var check = s

	for check != result {
		check = result
		result = ""

		var options = []string{}
		var length = len(s)
		var match = false

		for i := length - 1; i >= 0; i-- {
			if !unicode.IsDigit(rune(s[i])) {
				for o := len(options) - 1; o >= 0; o-- {
					option := options[o]
					match = true

					for j, m := range option {
						if (i-j) < 0 || rune(s[i-j]) != m {
							match = false
							break
						}
					}
					if match {
						count := 1
						olen := len(option)

						for i >= 0 {
							match = true

							for j, m := range option {
								if (i-j) < 0 || rune(s[i-j]) != m {
									match = false
									break
								}
							}
							if match {
								count++
								i -= olen
							} else {
								break
							}
						}
						rest := Reverse(options[0][:o])

						if olen == 1 {
							result = fmt.Sprintf("%d%s%s%s", count, option, rest, result)
						} else {
							result = fmt.Sprintf("%d(%s)%s%s", count, Reverse(option), rest, result)
						}
						match = true

						if i >= 0 && unicode.IsDigit(rune(s[i])) {
							count, err := strconv.Atoi(string(s[i]))
							if err != nil {
								return ""
							}
							result = strings.Repeat(string(s[i+1]), count-1) + result
						} else {
							i += 1
						}
						break
					}
				}
			}
			if match {
				options = []string{}
				match = false
			} else {
				s := string(s[i])

				for o := 0; o < len(options); o++ {
					options[o] = options[o] + s
				}
				options = append(options, s)
			}
		}
		if len(options) > 0 {
			result = Reverse(options[0]) + result
		}
		s = result
	}
	return result
}
