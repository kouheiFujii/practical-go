/*
構造体のフィールド名をJSON形式の文字列に変換する。
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
	u := User{
		Name: "John Doe",
		Age:  30,
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(b))
	// Output: {"Name":"John Doe","Age":30}
}
