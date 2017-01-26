package request

// セーブデータ送信リクエストEntity.
type SaveSend struct {
	// セーブデータ.
	SaveData string `json:"saveData"`
	// チェックデジット.
	CheckDigit string `json:"checkDigit"`
}

func (entity *SaveSend) Convert(
	object map[string]interface{},
) {
	if value, ok := object["saveData"].(string); ok {
		entity.SaveData = value
	}
	if value, ok := object["checkDigit"].(string); ok {
		entity.CheckDigit = value
	}
}
