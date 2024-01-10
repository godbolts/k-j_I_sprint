package sprint

func SplitWhitespaces(s string) []string {
	var words []string
	var word string

	for i, char := range s {
		if (string(char) != " " && string(char) != "\t" && string(char) != "\n") && i != len(s)-1 {
			word += string(char)
		} else {
			if i == len(s)-1 {
				word += string(char)
			}
			words = append(words, word)
			word = ""
		}
	}
	return words
}
