package sprint

func Overlap(arr1, arr2 []int) []int {
	var vastus []int
	counts := make(map[int]int)

	for _, v := range arr1 {
		counts[v]++
	}

	for _, v := range arr2 {
		if counts[v] > 0 {
			vastus = append(vastus, v)
			counts[v]--
		}
	}

	if len(vastus) == 0 {
		return []int{}
	}

	vastus = SortIntegerTable(vastus)

	return vastus
}

func SortIntegerTable(table []int) []int {
	for i := 0; i < len(table); i++ {
		min := i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[min] {
				min = j
			}
		}
		if min != i {
			table[i], table[min] = table[min], table[i]
		}
	}
	return table
}
