package sprint

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	s1 := strings.Split(str1, "")
	s2 := strings.Split(str2, "")

	sort.Strings(s1)
	sort.Strings(s2)

	return strings.Join(s1, "") == strings.Join(s2, "")
}
