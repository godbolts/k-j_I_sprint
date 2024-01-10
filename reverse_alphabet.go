package sprint

func ReverseAlphabet(step int) string {
	if step <= 0 {
		step = 1
	}
	var vastus string
	for ch := 'z'; ch >= 'a'; ch-- {
		vastus += string(ch)
		for i := 0; i < step-1; i++ {
			ch--
		}
	}
	return vastus
}
