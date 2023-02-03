/*
CSVファイルの読み込み方法
*/

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// 絶対パスを取得（アプリケーションの開始位置から指定する。/Users/kouhei/path/to/application/practical-go/）
	// absPath, _ := filepath.Abs("8: data-format/csv/country.csv")
	// CSVファイルを開く
	file, err := os.Open("8: data-format/csv/country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// CSVファイルを読み込む
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 各行を出力する
	for _, record := range records {
		fmt.Println(record)
	}
}
