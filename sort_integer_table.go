package sprint

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
