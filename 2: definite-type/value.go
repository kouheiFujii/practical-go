/*
SKUコード（Stock Keeping Unit, 在庫管理単位）を使用しているケース
T01230101のような値がクエリに付与されてくるパターン
*/
package main

import "fmt"

func main() {
	skuCD, _ := r.URL.Query("sku_code")
	// 部分文字列を指定して抽出する
	// コードの整合性チェックを呼び出し元で行わないといけないので少々やっかい
	itemCD, sizeCD, colorCD := skuCD[0:5], skuCD[5:7], skuCD[7:9]
}

/*
抽出する値に対して型を指定することで用途をわかりやすくする
*/

type SKUCode string

func (c SKUCode) Invalid() bool {
	// 桁数や利用可能文字のチェックを実施
	return false
}

func (c SKUCode) ItemCD() SKUCode {
	return c[0:5]
}

func (c SKUCode) SizeCD() SKUCode {
	return c[5:7]
}

func (c SKUCode) ColorCD() SKUCode {
	return c[7:9]
}

func main() {
	param := "T01230101"
	skuCD := SKUCode(param)

	if skuCD.Invalid() {
		// 異常系ハンドリング
	}
	itemCD, sizeCD, colorCD := skuCD.ItemCD(), skuCD.SizeCD(), skuCD.ColorCD()
	fmt.Printf("%v\n", itemCD)
	fmt.Println(sizeCD)
	fmt.Println(colorCD)
}
