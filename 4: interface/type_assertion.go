/*
型アサーションは、型を明示的に変換することを意味します。
Go言語では、「x.(T)」という形式で型アサーションを使用します。
xにはインターフェース型の変数、Tにはアサーションしたい型を記述します。

インターフェース越しに型にキャストして特定のフィールドにアクセスしたり機能を使用するときなどに使用される
*/

package main

import "fmt"

func main() {
	var x interface{} = "hello"
	s, ok := x.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("x is not a string")
	}
}
