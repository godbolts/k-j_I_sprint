package sprint

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
