/*
JSON 文字列が逐次的に入力される場合でも、デコードすることができる。
この方法では、io.Reader インタフェースを実装したデータソースから入力を取得することができる。

結論として、json.Unmarshal はメモリにすべてのデータが存在する場合に適していますが、
json.NewDecoder は、ストリームからデータを逐次的に読み込む場合に適しています。
*/

package main

import (
	"encoding/json"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// JSONデータを読み込む元となるio.Reader
	// ここでは、os.Stdinを使っています。
	reader := io.Reader(os.Stdin)

	// json.NewDecoderを使ってJSONデータを読み込む
	decoder := json.NewDecoder(reader)

	// JSONデータを格納する構造体を用意する
	var person Person

	// json.Decoder.Decodeを使ってJSONデータを読み込む
	err := decoder.Decode(&person)
	if err != nil {
		panic(err)
	}

	// 読み込んだ結果を表示する
	println(person.Name)
	println(person.Age)
}
