package sprint

type Point struct {
	X    float64
	Y    float64
	Text string
}

func MakePoint(x, y float64, text string) Point {
	return Point{
		X:    x,
		Y:    y,
		Text: text,
	}
}
