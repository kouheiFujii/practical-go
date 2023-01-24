/*
構造体への型定義
*/

package main

/*
アンチパターン

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

// 構造体をまるごと引数に取る
func ReadValue(r SensorData) float32 {
	if r.SensorType == "Fahrenheit" { // 華氏の場合は摂氏に変換
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}
*/

/*
関数の単体テストがしやすくなるので、レシーバを使用したほうがいい
構造体のフィールドを取り出して条件判定や計算をしている場合はレシーバにするチャンス
*/

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

func (r SensorData) ReadValue() float32 {
	if r.SensorType == "Fahrenheit" { // 華氏の場合は摂氏に変換
		return (r.Value * 9 / 5) + 32
	}
	return r.Value
}
