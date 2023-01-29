/*
インターフェースで定義されているメソッドと同名のメソッドを実装することにより、そのタイプはインターフェースを実装したことになる

特徴
1. 明示的な実装: インターフェースは明示的に実装する必要があります。型がインターフェースを暗黙的に実装しているわけではありません。

2. 型として使用できない: インターフェースは型として使用することができません（フィールドの定義ができない）。代わりに、構造体などの型のポインターがインターフェースに代入されます。

3. メソッド定義のみ: インターフェースはメソッドのシグネチャだけを定義します。実際の実装は実装する型が行います。

4. 多重実装: 型は複数のインターフェースを実装することができます。
*/

package main

import "fmt"

// Shapeインターフェース
type Shape interface {
	Area() float64
}

// Rectangle構造体
type Rectangle struct {
	width, height float64
}

// Circle構造体
type Circle struct {
	radius float64
}

// Rectangle構造体がShapeインターフェースを実装
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle構造体がShapeインターフェースを実装
func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func main() {
	// Rectangle構造体のポインター
	r := &Rectangle{width: 10, height: 5}

	// Circle構造体のポインター
	c := &Circle{radius: 5}

	// Shapeインターフェース型の配列
	shapes := []Shape{r, c}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}
