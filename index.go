package sprint

func Index(s string, toFind string) int {
	if len(s) == 0 && len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if i+len(toFind) <= len(s) && s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
