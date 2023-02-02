/*
Marshaler インターフェースを満たすことによりJSONエンコードの拡張が行える
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Marshaler インターフェースに実装されている関数
func (p Person) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"name":"%s","age":%d}`, p.Name, p.Age)), nil
}

func main() {
	person := Person{Name: "John Doe", Age: 30}

	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
	// Output: {"name":"John Doe","age":30}
}
