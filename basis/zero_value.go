/*
データの型に対応する初期値のことを「ゼロ値」という
変数に明示的に初期値を代入しない場合、その変数は対応するデータ型のゼロ値になる

・整数型: 0
・浮動小数点型: 0.0
・論理型: false
・文字列型: "" (空文字列)
・構造体型: 各フィールドのゼロ値
・スライス型: nil
・マップ型: nil
・ポインタ型: nil
・関数型: nil
・インターフェース型: nil
*/

package main

import "fmt"

func main() {
	var num1 int
	num2 := 0

	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	// num1: 0
	// num2: 0

	var str1 string
	str2 := ""

	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	// str1:
	// str2:
}
