package main

import "errors"

// https://go.dev/doc/effective_go
/*
・UrlでなくURLもしくはurl
・goでは短い変数が好まれる
・グローバルな変数にはわかり易い名前をローカル変数には短い名前を
*/
func main() {
	// MixedCapsを利用する
	var maxLength string

}

// Errorインターフェースを満たすError型は接尾頭にErrorをつける
type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

// エラー変数はErr
var (
	MarshalErr         = errors.New("marshal error")
	UpsupportedTypeErr = errors.New("upsupported type error")
	JsonErr            = errors.New("json error")
)
