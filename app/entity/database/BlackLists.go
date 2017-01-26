package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// ブラックリストテーブルEntity.
type BlackLists struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// ブラックユーザコード.
	BlackUserCode string `db:"black_user_code"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupBlackLists(db *gorm.DB) {
	db.AutoMigrate(&BlackLists{})
}
