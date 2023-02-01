/*
JSON形式の文字列を構造体に変換する方法です。
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	b := []byte(`{"Name":"John Doe","Age":30}`)

	var u User
	err := json.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(u)
	// Output: {John Doe 30}
}
