/*
バリエエーションを関数としてそれぞれ作成していく
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

// バリエーションが爆発するから麺の量は引数で渡す
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKituneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempulaUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}
