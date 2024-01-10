package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
	ascending := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			ascending = false
		}
	}
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if ascending {
			if f(arr[i], arr[i+1]) == 1 {
				return false
			}
		} else {
			if f(arr[i], arr[i+1]) == -1 {
				return false
			}
		}
	}
	return true
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
