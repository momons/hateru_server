package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

// アイテム交換テーブルEntity.
type ExchangeItems struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	UserCode string `db:"user_code"`
	// 交換トークン.
	ExchangeToken string `db:"exchange_token"`
	// ユーザ名.
	UserName string `db:"user_name"`
	// アイテム種類.
	ItemKindIndex int `db:"item_kind_index"`
	// アイテムコード.
	ItemCode string `db:"item_code"`
	// 希望アイテム種類.
	HopeItemKindIndex int `db:"hope_item_kind_index"`
	// 相手ユーザコード.
	PartnerUserCode string `db:"partner_user_code"`
	// 交換ステータス.
	ExchangeStatus string `db:"exchange_status"`
	// パスワードハッシュ.
	PasswordHash string `db:"password_hash"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// セットアップ.
func SetupExchangeItems(db *gorm.DB) {
	db.AutoMigrate(&ExchangeItems{})
}

// パートナーあり.
func (entity *ExchangeItems) IsHavePartner() bool {
	return len(entity.PartnerUserCode) > 0
}

// パスワードあり.
func (entity *ExchangeItems) HasPassword() bool {
	return len(entity.PasswordHash) > 0
}
