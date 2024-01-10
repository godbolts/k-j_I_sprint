package sprint

func AdvancedSortWordArr(a []string, f func(a, b string) int) []string {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if f(a[j], a[j+1]) == 1 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
