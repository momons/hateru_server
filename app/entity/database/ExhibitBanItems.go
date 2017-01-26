package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// アイテム交換禁止テーブルEntity.
type ExhibitBanItems struct {
	// ID.
	Id int64 `db:"id"`
	// アイテム種類.
	ItemKindIndex string `db:"item_kind_index"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupExhibitBanItems(db *gorm.DB) {
	db.AutoMigrate(&ExhibitBanItems{})
}
