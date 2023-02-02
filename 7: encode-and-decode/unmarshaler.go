/*
Unmarshaler インタフェースを利用したJSONデコードの拡張のサンプル
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func (u *User) UnmarshalJSON(b []byte) error {
	var obj map[string]interface{}
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}

	u.ID = int(obj["id"].(float64))
	u.Name = obj["name"].(string)
	u.Email = obj["email"].(string)
	u.IsAdmin = obj["is_admin"].(bool)

	return nil
}

func main() {
	data := []byte(`{"id": 1, "name": "John Doe", "email": "john@example.com", "is_admin": true}`)
	var user User
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}
