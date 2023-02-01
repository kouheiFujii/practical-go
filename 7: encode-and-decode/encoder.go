/*
逐次的にエンコードするからファイルの読み書きなどストリームを扱う処理のときメモリ効率が良い。

結論として、データ量が大きい場合や、生成したJSONデータを即座に利用する必要がある場合にはjson.NewEncoderを使用するのが良いです。
一方、全てのJSONデータを一度に生成し、後から利用する必要がある場合にはjson.Marshalを使用するのが良いです。
*/

package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "John Doe",
		Age:  30,
	}

	// JSONエンコーダを作成
	encoder := json.NewEncoder(os.Stdout)

	// データをJSONエンコーダに渡して逐次的にエンコードする
	err := encoder.Encode(p)
	if err != nil {
		// エラーハンドリング
		return
	}
}
