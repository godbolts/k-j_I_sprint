package sprint

import "strings"

func AtoiBase(s, base string) int {
	if len(base) < 2 {
		return 0
	}
	for _, char := range base {
		if char == '-' || char == '+' {
			return 0
		}
	}
	var res int
	for _, char := range s {
		index := strings.IndexRune(base, char)
		if index == -1 {
			return 0
		}
		res = res*len(base) + index
	}
	return res
}
