/*
インターフェースの結合をすることで複数のインターフェースを実装することができる
*/
package main

import "math"

type Shape interface {
	Area() float64
}

type Circle interface {
	Shape
	Circumference() float64
}

type Rectangle interface {
	Shape
	Perimeter() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return s.side * 4
}

func (s Square) Circumference() float64 {
	return s.side * 2 * math.Pi
}
