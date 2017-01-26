package request

// セーブデータ取得リクエストEntity.
type SaveGet struct {
	// セーブトークン.
	SaveToken string `json:"saveToken"`
}

func (entity *SaveGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["saveToken"].(string); ok {
		entity.SaveToken = value
	}
}
