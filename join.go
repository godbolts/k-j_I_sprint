package sprint

func Join(strs []string, sep string) string {
	var vastus string
	for i, word := range strs {
		vastus += word
		if i != len(strs)-1 {
			vastus += sep
		}
	}
	return vastus
}
