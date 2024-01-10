package sprint

func Sqrt(n int) int {
	if n < 0 {
		return 0
	}

	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return i
		}
	}
	return 0
}
