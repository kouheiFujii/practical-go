/*
ビルダーを利用したパターン

デメリット
コード量が増る。最後のOrder関数呼ばないと実行されない

メリット
コード補完が賢いエディタであればスムーズに実装できる
*/
package main

type Portion int

const (
	Regular Portion = iota + 1 // 普通
	Small                      // 小盛り
	Large                      //大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	// デフォルトはコンストラクタ関数で設定
	// 必須オプションはここに付与
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func useFluentInterfase() {
	oomorikitune := NewUdon(Large).Aburaage().Order()
}
