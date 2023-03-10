/*
Go言語において、クロージャは周囲のスコープから変数をキャプチャする無名関数です。クロージャを使用することで以下のことができます：

関数内で生成された変数の値を保存します（関数が完了しても）。
状態と振る舞いを関数内に封じ込めることで複雑な操作をシンプルにします。
動的な関数を生成することができます（実行時に関数の振る舞いを変えることができます）。
コールバック関数を実装することができます（別の関数に関数を引数として渡すことができます）。
複数回の呼び出しで状態を保持する必要がある場合や、複雑な操作をシンプルに行いたい場合、動的な関数を生成する必要がある場合などに、クロージャは有用です。状態と振る舞いを関数内に封じ込めることで、わかりやすく再利用可能なコードになります。
*/
package main

import "fmt"

func main() {
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
この例では、変数counterをキャプチャするクロージャincrementを作成します。各回、クロージャが呼び出されると、counterの値がインクリメントされ、その値が返されます。i
*/
