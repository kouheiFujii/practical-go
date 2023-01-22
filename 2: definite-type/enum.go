/*
列挙型のレシーバにロジックを実装する
*/
package main

import "fmt"

type Season int

const (
	Peak   Season = iota + 1 // 繁忙期
	Normal                   // 通常期
	Off                      // 閑散期
)

func (s Season) Price(price float64) float64 {
	// 繁忙期なら割増料金
	if s == Peak {
		return price + 200
	}
	return price
}

func main() {
	s := Peak.Price(10)
	fmt.Println(s)
	n := Normal.Price(100)
	fmt.Println(n)
}
