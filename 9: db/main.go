package main

// docker run -d --name my-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres

import (
	"context"
	"database/sql"
	"log"
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
	insert := `
		INSERT INTO users (user_id, user_name, created_at) VALUES ($1, $2, $3)
	`
	now := time.Now()

	_, err = db.Exec(insert, "001", "testuser1", now)
	if err != nil {
		log.Fatalf("insert error: %v", err)
	}

	_, err = db.Exec(insert, "002", "testuser2", now)
	if err != nil {
		log.Fatalf("insert error: %v", err)
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
