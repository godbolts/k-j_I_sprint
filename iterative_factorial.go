package sprint

func IterativeFactorial(n int) int {
	vastus := 1

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	for i := 1; i <= n; i++ {
		vastus *= i
	}
	return vastus
}
