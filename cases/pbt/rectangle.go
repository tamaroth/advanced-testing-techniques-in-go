package pbt

type Rectangle struct {
	Width, Height int
}

// Area returns the area of the rectangle.
func (r *Rectangle) Area() int {
	return r.Width * r.Height
}

// Resize changes the dimensions of the rectangle.
func (r *Rectangle) Resize(scaleFactor int) {
	r.Width *= scaleFactor
	r.Height *= scaleFactor
}

// Add adds the dimensions of another rectangle to the current rectangle.
func (r *Rectangle) Add(other *Rectangle) {
	r.Width += other.Width
	r.Height += other.Height
}
