package sprint

func Capitalize(s string) string {
	result := ""
	capNext := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capNext {
				if char >= 'a' && char <= 'z' {
					result += string(char - ('a' - 'A'))
				} else {
					result += string(char)
				}
				capNext = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + ('a' - 'A'))
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			capNext = true
		}
	}

	return result
}
