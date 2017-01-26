package request

// 掲示板書き込みリクエストEntity.
type BbsWrite struct {
	// 掲示板コード.
	BbsCode string `json:"bbsCode"`
	// メッセージコード.
	MessageCode string `json:"messageCode"`
	// メッセージタイプ.
	MessageType string `json:"messageType"`
	// メッセージデータ.
	MessageData string `json:"messageData"`
}

// コンバート.
func (entity *BbsWrite) Convert(
	object map[string]interface{},
) {
	if value, ok := object["bbsCode"].(string); ok {
		entity.BbsCode = value
	}
	if value, ok := object["messageCode"].(string); ok {
		entity.MessageCode = value
	}
	if value, ok := object["messageType"].(string); ok {
		entity.MessageType = value
	}
	if value, ok := object["messageData"].(string); ok {
		entity.MessageData = value
	}
}
