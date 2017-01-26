package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// アクセストークンテーブルEntity.
type AccessTokens struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// アクセストークン.
	AccessToken string `db:"access_token"`
	// 期限日時.
	PeriodAt time.Time `db:"period_at"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupAccessTokens(db *gorm.DB) {
	db.AutoMigrate(&AccessTokens{})
}
