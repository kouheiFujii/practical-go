package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
	_        // 使用されない(3が飛ばされる)
	d        // 4
	e        // 5
)

// ブロックごとでリセット
const (
	f = iota // 0
	g        // 1
	h        // 2
)

// べーすとなる型を使用して列挙型とする
type CarType int

const (
	// 初期化されていない0と見分けがつかないので1オリジンにするのが一般的
	Sedan CarType = iota + 1
	Hatchback
	MPV
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println("===============================")
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println("===============================")
	fmt.Println(Sedan)
	fmt.Println(Hatchback)
	fmt.Println(MPV)
	// error: cannot assign to Sedan (constant 1 of type CarType)
	// Sedan = 1
}
