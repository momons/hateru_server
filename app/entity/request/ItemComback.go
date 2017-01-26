package request

// アイテム取り戻しリクエストEntity.
type ItemComback struct {
	// 交換トークン.
	ExchangeToken string `json:"exchangeToken"`
}

// コンバート.
func (entity *ItemComback) Convert(
	object map[string]interface{},
) {
	if value, ok := object["exchangeToken"].(string); ok {
		entity.ExchangeToken = value
	}
}
