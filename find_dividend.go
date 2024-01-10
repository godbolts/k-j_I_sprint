package sprint

func FindDividend(from, to, divisor int) int {
	if divisor == 0 {
		return -1
	}
	for i := from; i < to; i++ {
		if i%divisor == 0 {
			return i
		} //
	}
	return -1
}
