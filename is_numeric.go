package sprint

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
