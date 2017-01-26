package request

// アイテム取り戻しリクエストEntity.
type ItemGet struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
	// 交換トークン.
	ExchangeToken string `json:"exchangeToken"`
	// 交換パスワード.
	Password string `json:"password"`
}

// コンバート.
func (entity *ItemGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["userCode"].(string); ok {
		entity.UserCode = value
	}
	if value, ok := object["exchangeToken"].(string); ok {
		entity.ExchangeToken = value
	}
	if value, ok := object["password"].(string); ok {
		entity.Password = value
	}
}
