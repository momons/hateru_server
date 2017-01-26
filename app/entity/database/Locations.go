package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 位置情報テーブルEntity.
type Locations struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// ユーザ名.
	UserName string `db:"user_name"`
	// メッセージ.
	Message string `db:"message"`
	// マップ.
	MapIndex int `db:"map_index"`
	// X座標.
	X int `db:"x"`
	// Y座標.
	Y int `db:"y"`
	// その他情報.
	OtherInfos string `db:"other_infos"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupLocations(db *gorm.DB) {
	db.AutoMigrate(&Locations{})
}
