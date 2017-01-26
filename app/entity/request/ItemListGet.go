package request

import "../../constants"

// アイテムリスト取得リクエストEntity.
type ItemListGet struct {
	// 交換タイプ.
	ExchangeType string `json:"exchangeType"`
	// アイテム種類インデックス.
	ItemKindIndex int `json:"itemKindIndex"`
	// 希望アイテム種類インデックス.
	HopeItemKindIndex int `json:"hopeItemKindIndex"`
	// 取得オフセット.
	Offset int `json:"offset"`
	// 取得数.
	Count int `json:"count"`
}

// コンバート.
func (entity *ItemListGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["exchangeType"].(string); ok {
		entity.ExchangeType = value
	}
	if value, ok := object["itemKindIndex"].(int); ok {
		entity.ItemKindIndex = value
	}
	if value, ok := object["hopeItemKindIndex"].(int); ok {
		entity.HopeItemKindIndex = value
	}
	if value, ok := object["offset"].(int); ok {
		entity.Offset = value
	}
	if value, ok := object["count"].(int); ok {
		entity.Count = value
	}
}

// 譲渡？
func (entity *ItemListGet) IsTransfer() bool {
	return entity.ExchangeType == constants.ExchangeTypeTransfer
}

// 交換？
func (entity *ItemListGet) IsExchange() bool {
	return entity.ExchangeType == constants.ExchangeTypeExchange
}
