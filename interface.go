// interface is a set of methods

package main

// that is interface
type Geometry interface {
	Area() float64
	Perimeter() float64
}

//that is struct
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
