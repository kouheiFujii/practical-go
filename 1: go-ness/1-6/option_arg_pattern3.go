/*
構造体を利用したパターン
標準ライブラリでよく見られる
ゼロ値対策が必要
*/
package main

import "time"

type Portion int

const (
	Regular Portion = iota + 1 // 普通
	Small                      // 小盛り
	Large                      //大盛り
)

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	// ゼロ値に対するデフォルト設定は関数内で処理する
	// 例：朝食時間は海老天1本無料
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}
