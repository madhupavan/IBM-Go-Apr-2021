package shapes

type Rect struct {
	Width, Height float32
}

func (r Rect) Area() float32 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float32 {
	return 2*r.Width + 2*r.Height
}
