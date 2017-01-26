package request

// キャラ取得リクエストEntity.
type CharaGet struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
}

// コンバート.
func (entity *CharaGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["userCode"].(string); ok {
		entity.UserCode = value
	}
}
