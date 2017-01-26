package request

// キャラ送信リクエストEntity.
type CharaSend struct {
	// ステータスデータ.
	StatusData string `json:"statusData"`
}

// コンバート.
func (entity *CharaSend) Convert(
	object map[string]interface{},
) {
	if value, ok := object["statusData"].(string); ok {
		entity.StatusData = value
	}
}
