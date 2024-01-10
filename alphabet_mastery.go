package sprint

func AlphabetMastery(n int) string {
	var vastus string
	for ch := 'a'; ch < 'a'+rune(n); ch++ {
		vastus += string(ch)
	}
	return vastus
}
