/*
型スイッチは、Go言語で型を検査する手法の一つです。
型スイッチは、特定の変数が特定の型であるかどうかを検査し、その型に応じて処理を実行することができます。
*/

package main

import "fmt"

func main() {
	var value interface{} = 7
	switch t := value.(type) {
	case int:
		fmt.Printf("int: %d\n", t)
	case float64:
		fmt.Printf("float64: %f\n", t)
	case string:
		fmt.Printf("string: %s\n", t)
	default:
		fmt.Printf("unexpected type\n")
	}
}
