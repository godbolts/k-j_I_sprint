package sprint

func ConvertBase(num, fromBase, toBase string) string {
	if len(fromBase) < 2 {
		return ""
	}
	for _, char := range fromBase {
		if char == '-' || char == '+' {
			return ""
		}
	}

	if len(toBase) < 2 {
		return ""
	}
	for _, char := range toBase {
		if char == '-' || char == '+' {
			return ""
		}
	}

	// rune counts
	fromLenRunes := StrLen(fromBase)
	toLenRunes := StrLen(toBase)
	numLen := StrLen(num)

	// loop over unicode runes in original string and store representative
	// values in number -- number[i] = index(num[i], fromBase)
	number, ipos := make([]int, numLen), 0
	for _, r := range num {
		jpos, found := 0, false
		for _, s := range fromBase {
			if r == s {
				number[ipos] = jpos
				found = true
				break
			}

			jpos++
		}

		// if character wasn't found in fromBase, then error
		if !found {
			return ""
		}

		ipos++
	}

	// split the runes in toBase
	todigits, idx := make([]rune, toLenRunes), 0
	for _, r := range toBase {
		todigits[idx] = r
		idx++
	}

	// loop until whole number is converted
	var result []rune
	for {
		divide, newlen := 0, 0

		// perform division manually (which is why this works with big numbers)
		for i := 0; i < numLen; i++ {
			divide = divide*fromLenRunes + number[i]
			if divide >= toLenRunes {
				number[newlen] = divide / toLenRunes
				divide = divide % toLenRunes
				newlen++
			} else if newlen > 0 {
				number[newlen] = 0
				newlen++
			}
		}

		numLen = newlen
		result = append(result, todigits[divide])

		if newlen == 0 {
			break
		}
	}

	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func StrLen(s string) int {
	vastus := len([]rune(s))
	return vastus
}
