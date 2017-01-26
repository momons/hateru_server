package manager

import (
	"../entity/database"
	"log"
)

// アイテム交換禁止マネージャ.
type ExhibitBanItems struct {
}

// インスタンス.
var instanceExhibitBanItems *ExhibitBanItems

// インスタンス取得.
func GetExhibitBanItems() *ExhibitBanItems {
	if instanceExhibitBanItems == nil {
		instanceExhibitBanItems = &ExhibitBanItems{}
	}
	return instanceExhibitBanItems
}

// アイテム種類インデックスで取得
func (manager *ExhibitBanItems) SelectOne(
	itemKindIndex int,
) *database.ExhibitBanItems {

	var entity database.ExhibitBanItems

	err := Db.Where(
		"item_kind_index = ?",
		itemKindIndex,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}
