package sprint

func LongestClimb(arr []int) []int {
	var longest []int
	var current []int

	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] > arr[i-1] {
			current = append(current, arr[i])
		} else {
			current = []int{arr[i]}
		}
		if len(current) > len(longest) {
			longest = current
		}
	}

	return longest
}
