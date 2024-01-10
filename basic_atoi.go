package sprint

func BasicAtoi(s string) int {
	var vastus int

	for i := 0; i < len(s); i++ {
		ch := s[i]
		arv := int(ch - '0')
		if arv >= 0 && arv <= 9 {
			vastus = vastus*10 + arv
		}
	}

	return vastus
}
