/*
スライスの場合、値を指定しないとnull値が設定される
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func main() {
	u1 := User{
		UserID:   "001",
		UserName: "gopher",
	}
	// nullにしたくない場合、明示的に空の配列を渡す必要がある
	u2 := User{
		UserID:    "002",
		UserName:  "gopher",
		Languages: []string{},
	}

	b1, _ := json.Marshal(u1)
	b2, _ := json.Marshal(u2)

	fmt.Println(string(b1))
	// {"user_id":"001","user_name":"gopher","languages":null}
	fmt.Println(string(b2))
	// {"user_id":"002","user_name":"gopher","languages":[]}
}
