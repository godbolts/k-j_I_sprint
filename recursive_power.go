package sprint

func RecursivePower(n int, power int) int {
	if power < 0 {
		return 0
	}

	if n < 0 {
		n = -n
	}

	if power == 0 {
		return 1
	} else {
		return n * RecursivePower(n, power-1)
	}
}
