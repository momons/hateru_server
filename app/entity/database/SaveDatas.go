package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// セーブデータテーブルEntity.
type SaveDatas struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// セーブトークン.
	SaveToken string `db:"save_token"`
	// セーブデータ
	SaveData string `db:"save_data"`
	// チェックデジット
	CheckDigit string `db:"check_digit"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupSaveDatas(db *gorm.DB) {
	db.AutoMigrate(&SaveDatas{})
}
