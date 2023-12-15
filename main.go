package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-bun-sample/user"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	// DSNを作成
	dsn := fmt.Sprintf(
		"%s:%s@tcp(localhost:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	// DBに接続
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// bunのDBインスタンスを作成
	db := bun.NewDB(sqldb, mysqldialect.New())

	// 検証用: 最初にテーブルを削除しておく
	db.NewDropTable().Model((*user.User)(nil)).Exec(context.Background())

	// usersテーブルを作成
	user.Create(context.Background(), db)

	u1 := &user.User{
		ID:   1,
		Name: "suzuki",
	}

	// レコードを作成
	u1.Insert(context.Background(), db)

	// レコードを取得
	u2 := &user.User{ID: 1}
	u2.FindByID(context.Background(), db)
	fmt.Println(u2.Name) // suzuki
}
