package main

// docker run -d --name my-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // ブランクインポートして初期化のみ実施
)

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

// テーブルの作成とデータの挿入
func createTable(db *sql.DB) error {
	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(32) NOT NULL,
			user_name VARCHAR(100) NOT NULL,
			created_at TIMESTAMP with time zone,
			CONSTRAINT pk_users PRIMARY KEY (user_id)
		);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		log.Fatalf("create table error: %v", err)
	}
	// テーブルにデータを挿入する
	// insert := `
	// 	INSERT INTO users (user_id, user_name, created_at) VALUES ($1, $2, $3)
	// `
	now := time.Now()
	/*
		パターン１:愚直に挿入

		_, err = db.Exec(insert, "001", "testuser1", now)
			if err != nil {
				log.Fatalf("insert error: %v", err)
			}

			_, err = db.Exec(insert, "002", "testuser2", now)
			if err != nil {
				log.Fatalf("insert error: %v", err)
			}
	*/
	/*
		パターン２：プリペアードステートメントを使って挿入
		プリペアードステートメントとは、SQLのテンプレートを作成しておき、あとは変数の値だけを指定して実行する。
		プリペアードステートメントを使うことで、SQLのパースやコンパイルを1回だけ行うことができる。

		users := []User{
			{UserID: "001", UserName: "testuser1", CreatedAt: now},
			{UserID: "002", UserName: "testuser2", CreatedAt: now},
			{UserID: "003", UserName: "testuser3", CreatedAt: now},
		}
		ctx := context.Background()
		stmt, err := db.PrepareContext(ctx, insert)
		if err != nil {
			log.Fatalf("prepare error: %v", err)
			return err
		}
		defer stmt.Close() // プリペアードステートメントはClose()で閉じる必要がある

		for _, user := range users {
			// 構築したプリペアードステートメントにパラメータをセットして実行する
			if _, err = stmt.ExecContext(ctx, user.UserID, user.UserName, user.CreatedAt); err != nil {
				log.Fatalf("exec error: %v", err)
				return err
			}
		}
	*/
	/*
		パターン３：バッチインサートを使用する
		バッチインサートとは、複数のINSERT文を1つのSQL文にまとめて実行することで、パフォーマンスを向上させる方法。
		例：
			INSERT INTO
				users (user_id, user_name, created_at)
			VALUES
				('001', 'testuser1', '2021-01-01 00:00:00'),
				('002', 'testuser2', '2021-01-01 00:00:00'),
				('003', 'testuser3', '2021-01-01 00:00:00');
	*/
	users := []User{
		{UserID: "001", UserName: "testuser1", CreatedAt: now},
		{UserID: "002", UserName: "testuser2", CreatedAt: now},
		{UserID: "003", UserName: "testuser3", CreatedAt: now},
	}
	// VALUESの部分を作成する
	valueStrings := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*3) // 3はカラム数
	number := 1
	for _, user := range users {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", number, number+1, number+2)) // $1, $2, $3
		valueArgs = append(valueArgs, user.UserID)                                                      // $1にはuser.UserIDが入る
		valueArgs = append(valueArgs, user.UserName)                                                    // $2にはuser.UserNameが入る
		valueArgs = append(valueArgs, user.CreatedAt)                                                   // $3にはuser.CreatedAtが入る
		number += 3                                                                                     // 3カラム文ずらす。$1, $2, $3の次は$4, $5, $6
	}
	query := fmt.Sprintf("INSERT INTO users (user_id, user_name, created_at) VALUES %s", strings.Join(valueStrings, ","))
	if _, err = db.Exec(query, valueArgs...); err != nil { // valueArgs...でスライスの展開
		log.Fatalf("exec error: %v", err)
		return err
	}

	return nil
}

// テーブルのデータを取得する(教科書のサンプル)
func selectTable(ctx context.Context, db *sql.DB) ([]*User, error) {
	selectQuery := `
		SELECT user_id, user_name, created_at FROM users ORDER BY user_id
	`
	rows, err := db.QueryContext(ctx, selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var (
			userID, userName string
			createdAt        time.Time
		)
		if err := rows.Scan(&userID, &userName, &createdAt); err != nil {
			return nil, err
		}
		users = append(users, &User{
			UserID:    userID,
			UserName:  userName,
			CreatedAt: createdAt,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("close error: %v", err)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("rows error: %v", err)
	}
	return users, nil
}

func main() {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		log.Fatalf("db err: %v", err)
		panic("failed db setup")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed pin: %v", err)
	}

	// テーブルが存在するか確認
	existsTable := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_schema = 'public'
			AND table_name = 'users'
		);
	`
	var exists bool
	err = db.QueryRow(existsTable).Scan(&exists)
	if err != nil {
		log.Fatalf("exists table error: %v", err)
	}

	if !exists {
		// テーブルが存在しない場合はテーブルを作成する
		if err := createTable(db); err != nil {
			log.Fatalf("create table error: %v", err)
		}
	}

	// テーブルのデータを取得する
	ctx := context.Background()
	users, err := selectTable(ctx, db)
	if err != nil {
		log.Fatalf("select error: %v", err)
	}
	for _, user := range users {
		log.Printf("user: %v", user)
	}
}
