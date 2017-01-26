package request

// アイテム情報取得リクエストEntity.
type ItemInfoGet struct {
	// ユーザコード
	UserCode string `json:"userCode"`
	// 交換トークン
	ExchangeToken string `json:"exchangeToken"`
}

// コンバート.
func (entity *ItemInfoGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["userCode"].(string); ok {
		entity.UserCode = value
	}
	if value, ok := object["exchangeToken"].(string); ok {
		entity.ExchangeToken = value
	}
}
