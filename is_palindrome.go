package sprint

func IsPalindrome(s string) bool {
	s = ToLower(s)
	round := ""
	for i := len(s) - 1; i >= 0; i-- {
		round += string(s[i])
	}
	if s == round {
		return true
	} else {
		return false
	}

}

func ToLower(s string) string {
	var vastus string
	for i := 0; i < len(s); i++ {

		roll := s[i]
		if roll >= 'A' && roll <= 'Z' {
			roll = roll + 32

		}
		vastus += string(roll)
	}
	return vastus

}
