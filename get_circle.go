package sprint

type Circle struct {
	Radius    float64
	Diameter  float64
	Area      float64
	Perimeter float64
}

func GetCircle(r float64) Circle {
	π := 3.14
	return Circle{
		Radius:    r,
		Diameter:  2 * r,
		Area:      π * r * r,
		Perimeter: 2 * π * r,
	}
}
