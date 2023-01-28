package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

/*
値レシーバー: レシーバーのコピーを操作するメソッド。変数のポインタを使用することはできない。
*/

func (p Person) addAge() {
	p.Age++
}

/*
ポインタレシーバー: レシーバーのポインタを操作するメソッド。変数のコピーを操作することはできない。
*/

func (p *Person) addAgePointer() {
	p.Age++
}

/*
値レシーバーはメソッド内で変数を操作しても、呼び出し元に影響を与えません。
一方、ポインタレシーバーはメソッド内で変数を操作することで、呼び出し元も変更されます。

呼び出し元で引き続き使用したい場合や、大きなオブジェクトを扱う場合はポインタレシーバーを使用するとメモリ効率が良くなりますが、
小さなオブジェクトを扱う場合や、メソッド内での変数の操作によって呼び出し元に影響を与えたくない場合は値レシーバーを使用するのが望ましいと言えます。

小さなオブジェクトを扱う場合のメモリ効率については使用する環境は要件によって変わってくる。
*/

func main() {
	p := Person{"John", 20}

	fmt.Println("before addAge:", p)
	p.addAge()
	fmt.Println("after addAge:", p)

	fmt.Println("before addAgePointer:", p)
	p.addAgePointer()
	fmt.Println("after addAgePointer:", p)

	/*
		before addAge: {John 20}
		after addAge: {John 20}
		before addAgePointer: {John 20}
		after addAgePointer: {John 21}
	*/
}
