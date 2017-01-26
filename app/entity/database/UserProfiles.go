package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// プロフィールテーブルEntity.
type UserProfiles struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// ユーザ名.
	UserName string `db:"user_name"`
	// プロフィールデータ
	ProfileData string `db:"profile_data"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupUserProfiles(db *gorm.DB) {
	db.AutoMigrate(&UserProfiles{})
}
