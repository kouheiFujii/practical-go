package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"書籍名", "出版年", "ページ数"},
		{"Go言語によるWebアプリケーション開発", "2016", "280"},
		{"Go言語による並行処理", "2018", "256"},
		{"Go言語でつくるインタプリタ", "2018", "316"},
	}
	// ディレクトリを指定して生成することも可能 /path/to/directory/oreilly.csv
	file, err := os.OpenFile("oreilly.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Cannot create file: %s\n", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush() // バッファリングされているデータの書き込み

	for _, record := range records {
		err := writer.Write(record) // 実際に書き込まれているわけでなくバッファリングに蓄積されている
		if err != nil {
			log.Fatalf("Cannot write to file: %s\n", err)
		}
	}

}
