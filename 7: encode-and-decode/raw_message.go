/*
json.RawMessageはJSONの値を一旦そのまま保持することができる
フィールドの状態に応じて異なる構造体を切り替えることができる
*/

package main

import (
	"encoding/json"
	"fmt"
)

type wrapper struct {
	Kind string
	Data json.RawMessage
}

type data1 struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

type data2 struct {
	Baz string `json:"baz"`
	Qux int    `json:"qux"`
}

func main() {
	// JSON 入力
	input := []byte(`{"kind": "data1", "data": {"foo": "hello", "bar": 42}}`)

	// JSON の受け皿となる wrapper 構造体を作成
	var w wrapper
	if err := json.Unmarshal(input, &w); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("wrapper:", w)
	// wrapper: {data1 [123 34 102 111 111 34 58 32 34 104 101 108 108 111 34 44 32 34 98 97 114 34 58 32 52 50 125]}

	// Kind に応じて適切な構造体にデコードする
	switch w.Kind {
	case "data1":
		var d1 data1
		if err := json.Unmarshal(w.Data, &d1); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("data1:", d1)
		// data1: {hello 42}
	case "data2":
		var d2 data2
		if err := json.Unmarshal(w.Data, &d2); err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println("data2:", d2)
	}
}
