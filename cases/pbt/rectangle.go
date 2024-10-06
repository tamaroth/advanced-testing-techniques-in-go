package pbt

type Rectangle struct {
	Width, Height int
}

// Area returns the area of the rectangle.
func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of the rectangle.
func (r *Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

// Resize changes the dimensions of the rectangle.
func (r *Rectangle) Resize(scaleFactor int) {
	r.Width *= scaleFactor
	r.Height *= scaleFactor
}
