/*
defer の落とし穴
関数を抜けるまで実行されないので、forループの中で使用すると線形的にリソースを使用する
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {

for _, fi := range files {
	f, err :=os.Open(file)
	if err != nil {
		return err
	}
	// この書き方はforループを抜けるまでdeferが実行されない
	// deferを使わずにファイルの使用後に自分で f.Close を呼ぶ
	defer f.Close()
	dara, _ := io.ReadAll(f)
	result = append(result, data)
}
}

// エラー処理系のメソッドは普通にdefer呼び出すだけではエラーを取りこぼす
// 無名関数にしてそのエラーを名前付き返り値に代入すると呼び出し元に返せる
func deferReturnSample(fname string) (defErr error) {
	var f *os.File
	f, err :=os.Create(fname)
	if err!=nil {
		return fmt.Errorf("failed open file: %v", err)
	}
	defer func ()  {
		// Closeのエラーを拾って名前付き返り値に代入
		defErr = f.Close()
	}()
	io.WriteString(f, "defer return sample")
	return
}
