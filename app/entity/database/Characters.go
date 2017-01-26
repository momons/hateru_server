package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// キャラテーブルEntity.
type Characters struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// ユーザ名.
	UserName string `db:"user_name"`
	// ステータスデータ.
	StatusData string `db:"status_data"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupCharacters(db *gorm.DB) {
	db.AutoMigrate(&Characters{})
}
