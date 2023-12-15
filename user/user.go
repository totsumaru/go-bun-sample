package user

import (
	"context"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:"id,pk,autoincrement"`
	Name  string `bun:"name,notnull"`
	email string // unexported fields are ignored
}

// Create: テーブルを作成します
//
// すでにテーブルが存在する場合はエラーを返します。
func Create(ctx context.Context, db *bun.DB) {
	_, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

// Insert: レコードを挿入します
func (u *User) Insert(ctx context.Context, db *bun.DB) {
	_, err := db.NewInsert().Model(u).Exec(ctx)
	if err != nil {
		panic(err)
	}
}

// Select: IDでレコードを取得します
func (u *User) FindByID(ctx context.Context, db *bun.DB) {
	if err := db.NewSelect().Model(u).Where("id = ?", u.ID).Scan(ctx); err != nil {
		panic(err)
	}
}
