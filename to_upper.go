package sprint

func ToUpper(s string) string {
	var vastus string
	for i := 0; i < len(s); i++ {

		roll := s[i]
		if roll >= 'a' && roll <= 'z' {
			roll = roll - 32

		}
		vastus += string(roll)
	}
	return vastus

}
