package dto

type Shape interface {
	Perimeter() float64
	Area() float64
	GetType() string
}
type Rectangle struct {
	Width  float64
	Height float64
}
type Square struct {
	Width float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return (rectangle.Width + rectangle.Height) * 2
}
func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}
func (square Square) Perimeter() float64 {
	return square.Width * 4
}
func (square Square) Area() float64 {
	return square.Width * 2
}
func (square Square) GetType() string {
	return "square"
}
func (rectangle Rectangle) GetType() string {
	return "rectangle"
}
