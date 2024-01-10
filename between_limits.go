package sprint

func BetweenLimits(from, to rune) string {
	if from > to {
		from, to = to, from
	}
	var result string

	for ch := from + 1; ch < to; ch++ {
		result += string(ch)
	}

	return result
}
