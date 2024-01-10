package sprint

func Accumulate(n int) int {
	if n <= 0 {
		return 0
	}
	korrutis := 0
	for i := 1; i <= n; i++ {
		korrutis += i
	}
	return korrutis
}
