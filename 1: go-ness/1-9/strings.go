/*
Goの文字列は不変

すでにある文字列に追加や取り出しなどといった加工を行うとその都度新たにメモリを確保して文字列を生成する
*/
package main

import (
	"log"
	"strings"
)

func main() {
	// Bad pattern
	// その都度メモリが確保されて処理が遅くなる
	src := []string{"Back", "To", "The", "Future", "Part", "Ⅲ"}
	var title string
	for i, word := range src {
		if i != 0 {
			title += " "
		}
		titie += word
	}
	log.Println(title)

	// Good pattern
	var builder strings.Builder
	// 最終的な文字列がわかっているならば Grow を利用するとより良い
	builder.Grow(100) // 100文字以下と仮定
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())

	// 数が少ないのであれば一つの式にまとめてしまうのも悪くない
	displayTitle := "1990年7月6日 - " + title + " - ロバート・ゼメキス"
	log.Println(displayTitle)
}
