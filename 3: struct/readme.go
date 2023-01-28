package main

import (
	"fmt"
	"time"
)

/*
構造体の基本
・他の構造体を埋め込むことができる。オブジェクト指向の継承とは少し違ってダウンキャストやアップキャストは出来ない
・関数や式の中でも動的に定義することができる
*/

// 小文字ならプライベート。大文字ならパブリック
// 他の構造体を埋め込むことできる
type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title      string // フィールド
	Author     Author
	Pubulisher string
	ReleasedAt time.Time
	ISBN       string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	// 構造体を利用してインスタンスを作成する
	author := Author{
		FirstName: "渋川", // フィールドに対応する値を書いていることを複合リテラルという
		LastName:  "よしき",
	}
	book := Book{
		Title:      "Real World HTTP",
		Author:     author,
		Pubulisher: "オライリー・ジャパン",
		ISBN:       "48731190",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book.Title)

	// 一時的な構造体を用意してすぐにインスタンス化する
	book2 := struct {
		Title      string
		Author     string
		Pubulisher string
		ReleasedAt time.Time
		ISBN       string
	}{
		Title:      "Real World HTTP",
		Author:     "渋川よしき",
		Pubulisher: "オライリー・ジャパン",
		ISBN:       "48731190",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book2.Author)

	/*
		Tips
		構造体の裏ではポインタが使われている
		ピリオドを使用したメンバアクセスで、自動的にデリファンス（ポインタからの値取り出し）を行う
		格納されているのがポインタのポインタとかの場合は自分でデリファンスして値を取り出す必要がある。

		メモ
		go version go1.18.6 では 以下のやり方しなくても取り出せた。バージョンアップによって改善してるのかも
	*/
	book3 := Book{
		Title: "Mithril",
	}
	fmt.Println(book3.Title) // OK
	fmt.Println((*&book3).Title)

	book4 := &book3
	fmt.Println(book4.Title) // NG(go version go1.18.6 では出力される)
	fmt.Println((**&book4).Title)
}
