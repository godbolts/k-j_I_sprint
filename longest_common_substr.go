package sprint

func LongestCommonSubstr(str1, str2 string) string {

	m := make([][]int, len(str1))
	longest := 0
	end := 0

	for i := range m {
		m[i] = make([]int, len(str2))
		for j := range m[i] {
			if str1[i] == str2[j] {
				if i == 0 || j == 0 {
					m[i][j] = 1
				} else {
					m[i][j] = m[i-1][j-1] + 1
				}
				if m[i][j] > longest {
					longest = m[i][j]
					end = i + 1
				}
			}
		}
	}

	return str1[end-longest : end]
}
