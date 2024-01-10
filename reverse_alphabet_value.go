package sprint

func ReverseAlphabetValue(ch rune) rune {
	shift := int(ch) % 97
	newRune := 122 - rune(shift)
	return newRune
}
