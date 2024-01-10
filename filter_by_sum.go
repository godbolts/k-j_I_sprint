package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
	var vastus [][]int
	for i := 0; i < len(arr); i++ {
		sum := 0
		for _, num := range arr[i] {
			sum += num
		}
		if sum >= limit {
			vastus = append(vastus, arr[i])
		}
	}
	if len(vastus) == 0 {
		return [][]int{}
	}
	return vastus
}
