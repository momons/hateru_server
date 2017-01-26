package request

// 位置情報送信リクエストEntity.
type LocationSend struct {
	// メッセージ.
	Message string `json:"message"`
	// マップインデックス.
	MapIndex int `json:"mapIndex"`
	// X座標.
	X int `json:"x"`
	// Y座標.
	Y int `json:"y"`
	// その他情報
	OtherInfos string `json:"otherInfos"`
}

func (entity *LocationSend) Convert(
	object map[string]interface{},
) {
	if value, ok := object["message"].(string); ok {
		entity.Message = value
	}
	if value, ok := object["mapIndex"].(int); ok {
		entity.MapIndex = value
	}
	if value, ok := object["x"].(int); ok {
		entity.X = value
	}
	if value, ok := object["y"].(int); ok {
		entity.Y = value
	}
	if value, ok := object["otherInfos"].(string); ok {
		entity.OtherInfos = value
	}
}
