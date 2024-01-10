package sprint

import (
	"strconv"
	"strings"
)

var solutions []string

func placeQueens(queens [8]int, ob int) {
	for i := 0; i < 8; i++ {
		if canPlace(queens, ob, i) {
			queens[ob] = i
			if ob == 7 {
				addSolution(queens)
			} else {
				placeQueens(queens, ob+1)
			}
		}
	}
}

func canPlace(queens [8]int, ocuppiedRows, numberToCheck int) bool {
	for i := 0; i < ocuppiedRows; i++ {
		if queens[i] == numberToCheck || queens[i]-i == numberToCheck-ocuppiedRows || queens[i]+i == numberToCheck+ocuppiedRows {
			return false
		}
	}
	return true
}

func addSolution(queens [8]int) {
	var sb strings.Builder
	for _, queen := range queens {
		sb.WriteString(strconv.Itoa(queen + 1))
	}
	solutions = append(solutions, sb.String())
}

func EightQueens() string {
	solutions = []string{}
	queens := [8]int{}
	placeQueens(queens, 0)
	return strings.Join(solutions, "\n")
}
