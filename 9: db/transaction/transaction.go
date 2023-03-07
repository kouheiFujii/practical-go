package transaction

import (
	"context"
	"database/sql"
	"fmt"
)

/*
シンプルなトランザクションパターン
エラーが発生したらその都度ロールバックする
*/
// type Service struct {
// 	db *sql.DB
// }

// func (s *Service) UpdateUser(ctx context.Context, userID string) error {
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	updateQuery := `
// 		UPDATE users SET user_name = $1 WHERE user_id = $2
// 	`
// 	if _, err := tx.ExecContext(ctx, updateQuery, "updated", userID); err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return tx.Commit()
// }

/*
defer を使ってロールバックする
トランザクション処理中にエラーや panic が発生したら確実にロールバックすることを保証する
*/

// type Service struct {
// 	db *sql.DB
// }

// func (s *Service) UpdateUser(ctx context.Context, userID string) error {
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback() // トランザクション処理中にエラーが発生したら確実にロールバックする

// 	updateQuery := `
// 		UPDATE users SET user_name = $1 WHERE user_id = $2
// 	`
// 	if _, err := tx.ExecContext(ctx, updateQuery, "updated", userID); err != nil {
// 		return err
// 	}
// 	return tx.Commit()
// }

/*
トランザクションのラッパーを作成してトランザクション処理と実装の分離を行う
*/

// トランザクションを制御するための構造体
type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx *txAdmin
}

// トランザクションを制御するメソッド
// 開発者は本メソッドを使用してDMLクエリを実行する
func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := f(ctx); err != nil {
		return fmt.Errorf("transaction query failed: %w", err)
	}
	return tx.Commit()
}

// 使用方法
func (s *Service) UpdateUser(ctx context.Context, userID string) error {
	updateQuery := `
		UPDATE users SET user_name = $1 WHERE user_id = $2
	`
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, updateQuery, "updated", userID); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}
