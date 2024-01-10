package sprint

func Split(s, sep string) []string {
	var words []string
	var word string
	var i int

	if len(s) == 0 {
		return []string{}
	}

	for i < len(s) {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			words = append(words, word)
			word = ""
			i += len(sep)
		} else {
			word += string(s[i])
			i++
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}
