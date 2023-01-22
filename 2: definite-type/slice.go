/*
スライスに対して定義型を作成。
データベースなどへの問い合わせ結果をレシーバに移譲する
ユニットテストがしやすくなる
*/

package main

import "time"

type Consumer struct {
	ActiveFlg bool
}

type Consumers []Consumer

func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlg {
			resp = append(resp, v)
		}
	}
	return resp
}

/*
利用例
gets, err := GetConsumers(ctx, key)
activeConsumers := gets.ActiveConsumer()
*/

// ロジックがコレクション操作なら戻り値も揃えておくと良い
// 複数のコレクション操作を組み合わせたい場合に、メソッドチェーンで記述できるようになる
// consumers := gets.ActiveConsumer().Expires(time.Now().AddDate(0, 1, 0)).SortByExpiredAt()

// メソッドチェーンを兼ね備えた関数を宣言することも可能
// 関数を他パッケージに対してプライベートにすることで実装をカプセル化してる
func (c Consumers) RequiredFollows() Consumers {
	return c.activeConsumers().expires(time.Now().AddDate(0, 1, 0)).sortByExpiredAt()
}

func (c Consumers) activeConsumers() Consumers {
	// 契約が有効なユーザの絞り込み
}

func (c Consumers) expires(end time.Time) Consumers {
	// endの日時以降に契約が執行するユーザに絞り込み
}

func (c Consumers) sortByExpiredAt() Consumers {
	// 契約期限日で昇順にソート
}
