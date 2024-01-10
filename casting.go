package sprint

import "math"

func Casting(n float64) int {
	rounded := int(math.Round(n))
	return rounded
}
