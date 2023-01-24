/*
機密情報やユーザのクレジットカードの番号など誤ってlog出力されたときに情報を出力したくないケース
Stringerを利用してマスキングする
*/
package main

import (
	"encoding/json"
	"fmt"
)

type CreditCard string

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

// 基本のStringer
func (c ConfidentialCustomer) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

// fmtの書式指定子を %#v にすると出力できてしまうのでこちらもカバー
func (c ConfidentialCustomer) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func main() {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1111-1111-1111",
	}

	fmt.Println(c)         // xxxx-xxxx-xxxx-xxxx
	fmt.Printf("%v\n", c)  // xxxx-xxxx-xxxx-xxxx
	fmt.Printf("%+v\n", c) // xxxx-xxxx-xxxx-xxxx
	fmt.Printf("%#v\n", c) // xxxx-xxxx-xxxx-xxxx

	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes)) // もとの値
	// JSON:  {"CustomerID":1,"CreditCard":"4111-1111-1111-1111"}
}
