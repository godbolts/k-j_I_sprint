package sprint

func IterativePower(n int, power int) int {
	vastus := 1

	if power < 0 {
		return 0
	}

	if n < 0 {
		n = -n
	}

	for i := 0; i < power; i++ {
		vastus *= n
	}

	return vastus
}
