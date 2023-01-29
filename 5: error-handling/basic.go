/*
基本のエラーハンドリング用の型
*/

package main

import "fmt"

type StatusCode int

const (
	Success             StatusCode = 200
	BadRequest                     = 400
	Unauthorized                   = 401
	Forbidden                      = 403
	NotFound                       = 404
	InternalServerError            = 500
)

// 独自のエラー型
type MyError struct {
	Message string
	Code    StatusCode
}

// error インターフェースの実装
func (e MyError) Error() string {
	return e.Message
}

func doSomething() error {
	// some operation
	if someCondition {
		return MyError{"an error occurred", BadRequest}
	}
	return nil
}

func main() {
	if err := doSomething(); err != nil {
		// 型アサーションでエラー結果を出力する
		if myError, isMyError := err.(MyError); isMyError {
			fmt.Println("Error Message:", myError.Message)
			fmt.Println("Error Code:", myError.Code)
		}
	}
}
