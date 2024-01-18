package structsmethodinterfaces


func Perimeter(rectangle Rectrangle) float64{
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectrangle) float64 {
	return rectangle.Width * rectangle.Height
}