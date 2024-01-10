package sprint

func MakeRange(min, max int) []int {
	var vastus []int
	if min >= max {
		return vastus
	}

	for i := 0; i < (max - min); i++ {
		vastus = append(vastus, min+i)
	}
	return vastus
}
