// package main

// import "fmt"

// type Calulate interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Square struct {
// 	Length float64
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// type Circle struct {
// 	Radius float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// func (r Rectangle) Perimeter() float64 {
// 	return 2 * (r.Width + r.Height)
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func (c Circle) Perimeter() float64 {
// 	return 2 * 3.14 * c.Radius
// }

// func (s Square) Area() float64 {
// 	return s.Length * s.Length
// }

// func (s Square) Perimeter() float64 {
// 	return 4 * s.Length
// }

// func main() {
// 	var rectangle Calulate = Rectangle{10, 20}
// 	fmt.Println("Area of rectangle is ", rectangle.Area())
// 	fmt.Println("Perimeter of rectangle is ", rectangle.Perimeter())

// 	var circle Calulate = Circle{10}
// 	fmt.Println("Area of circle is ", circle.Area())
// 	fmt.Println("Perimeter of circle is ", circle.Perimeter())

// 	var square Calulate = Square{10}
// 	fmt.Println("Area of square is ", square.Area())
// 	fmt.Println("Perimeter of square is ", square.Perimeter())

// }
