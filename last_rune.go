package sprint

func LastRune(s string) rune {
	arv := len(s) - 1
	b_element := s[arv]
	return rune(b_element)
}
