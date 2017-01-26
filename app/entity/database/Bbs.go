package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 掲示板テーブルEntity.
type Bbs struct {
	// ID.
	Id int64 `db:"id"`
	// 掲示板コード.
	BbsCode string `db:"bbs_code"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// ユーザ名.
	UserName string `db:"user_name"`
	// メッセージコード.
	MessageCode string `db:"message_code"`
	// メッセージタイプ.
	MessageType string `db:"message_type"`
	// メッセージデータ.
	MessageData string `db:"message_data"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupBbs(db *gorm.DB) {
	db.AutoMigrate(&Bbs{})
}
