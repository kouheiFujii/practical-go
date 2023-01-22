/*
Functional Option パターン

メリデメはビルダーパターンと似ている
返り値の関数を型定義しておくと GoDoc に反映できる

デメリット
コードが横に長くなりがち

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

type OptFunc func(r *Udon)

func NewUdon(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.men = p }
}

func OptAburaage() OptFunc {
	return func(r *Udon) { r.aburaage = true }
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) { r.ebiten = n }
}

func useFuncOption() {
	tokuseiUdon := NewUdon(OptAburaage(), OptEbiten(3))
}
