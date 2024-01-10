package sprint

func RemoveDuplicates(arr []int) []int {
	var vastus []int

	korrad_a := make(map[int]int)
	for _, char := range arr {
		korrad_a[char]++
		if korrad_a[char] > 1 {
			continue
		} else {
			vastus = append(vastus, char)
		}
	}
	return vastus
}
