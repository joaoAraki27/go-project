package game

const (
	screenWidth  = 1200
	screenHeight = 920
)

type Vector struct {
	x float64
	y float64
}

type Rect struct {
	x      float64
	y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{
		x:      x,
		y:      y,
		Width:  width,
		Height: height,
	}
}

func (r Rect) Intersects(other Rect) bool {
	return r.x <= other.maxX() &&
		other.x <= r.maxX() &&
		r.y <= other.maxY() &&
		other.y <= r.maxY()
}

func (r Rect) maxX() float64 {
	return r.x + r.Width
}

func (r Rect) maxY() float64 {
	return r.y + r.Height
}
