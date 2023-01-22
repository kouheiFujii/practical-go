package main

import "fmt"

type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

// 文字列でステータステキストを返す
func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauth0rized"
	case StatusPaymentRequired:
		return "PaymentReqired"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}

// 日本の国道を返す型
type NationalRoute int

const (
	NagasakiKaido   NationalRoute = 200
	AizuNumataKaido NationalRoute = 401
	HokurikuDo      NationalRoute = 402
	KurinokiByhpass NationalRoute = 403
)

// 同じString メソッドを持つが型で判断するのでエラーになることはない
func (n NationalRoute) String() string {
	switch n {
	case NagasakiKaido:
		return "長崎街道"
	case AizuNumataKaido:
		return "会津沼田街道"
	case HokurikuDo:
		return "北陸道"
	case KurinokiByhpass:
		return "栗の木バイパス"
	default:
		return fmt.Sprintf("国道%d号線", n)
	}
}
