package sprint

func Atoi(s string) int {
	var vastus int
	var neg bool
	invalidSymbolCount := 0
	if s[0] == '-' && s[1] != '-' {
		neg = true
	}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		arv := int(ch - '0')
		if arv >= 0 && arv <= 9 {
			vastus = vastus*10 + arv
		} else {
			invalidSymbolCount++
		}
	}

	if invalidSymbolCount >= 2 {
		return 0
	}
	if neg {
		vastus = -vastus
	}
	return vastus
}
