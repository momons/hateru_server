package request

// アイテム送信リクエストEntity.
type ItemSend struct {
	// アイテム種類.
	ItemKindIndex int `json:"itemKindIndex"`
	// アイテム種類.
	ItemCode string `json:"itemCode"`
	// 交換希望アイテム種別.
	HopeItemKindIndex int `json:"hopeItemKindIndex"`
	// 交換相手ユーザコード.
	PartnerUserCode string `json:"partnerUserCode"`
	// 交換パスワード.
	Password string `json:"password"`
}

// コンバート.
func (entity *ItemSend) Convert(
	object map[string]interface{},
) {
	if value, ok := object["itemKindIndex"].(int); ok {
		entity.ItemKindIndex = value
	}
	if value, ok := object["itemCode"].(string); ok {
		entity.ItemCode = value
	}
	if value, ok := object["hopeItemKindIndex"].(int); ok {
		entity.HopeItemKindIndex = value
	}
	if value, ok := object["partnerUserCode"].(string); ok {
		entity.PartnerUserCode = value
	}
	if value, ok := object["password"].(string); ok {
		entity.Password = value
	}
}
