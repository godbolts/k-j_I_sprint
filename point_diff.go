package sprint

type Point struct {
	X    float64
	Y    float64
	Text string
}

func PointDiff(p1, p2 Point) Point {
	if p1.X > p2.X || p1.Y > p2.Y {
		return p1
	}
	return p2
}
