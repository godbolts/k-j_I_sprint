package sprint

func CountDivisible(from, to, step, divisor int) int {
	if step <= 0 {
		return 0
	}
	if divisor == 0 {
		return 0
	}

	var kokku int
	for i := from; i <= to; i += step {
		if i%divisor == 0 {
			kokku += 1
		}
	}
	return kokku
}
