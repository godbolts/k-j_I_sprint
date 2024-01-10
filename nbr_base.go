package sprint

func NbrBase(num int, base string) string {
	var vastus string
	var negatiivne bool
	var korduvad bool
	if len(base) < 2 {
		return "NV"
	}
	korrad := make(map[rune]int)
	for _, char := range base {
		korrad[char]++
		if korrad[char] > 1 {
			korduvad = true
		}

	}

	if korduvad {
		return "NV"
	}

	for _, char := range base {
		if char == '-' || char == '+' {
			return "NV"
		}
	}
	if num < 0 {
		num = -num
		negatiivne = true
	}

	baseLen := len(base)

	var result []rune
	for num > 0 {
		remainder := num % baseLen
		result = append(result, rune(base[remainder]))
		num = num / baseLen
	}

	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	vastus = string(result)

	if negatiivne {
		vastus = "-" + vastus
	}

	return vastus
}
