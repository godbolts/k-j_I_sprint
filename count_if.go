package sprint

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for _, element := range tab {
		if f(element) {
			count++
		}
	}
	return count
}

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsAlphanumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		char := s[i]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	lower := true

	for i := 0; i < len(s); i++ {
		element := s[i]

		if element < 'a' || element > 'z' {
			lower = false
		}
	}
	return lower
}

func IsNumeric(s string) bool {
	vastus := true

	for i := 0; i < len(s); i++ {
		roll := s[i]
		if roll < '0' || roll > '9' {
			vastus = false
		}
	}

	return vastus
}
