package request

// 位置情報取得リクエストEntity.
type LocationGet struct {
	// マップインデックス.
	MapIndex int `json:"mapIndex"`
	// X座標.
	X int `json:"x"`
	// Y座標.
	Y int `json:"y"`
}

func (entity *LocationGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["mapIndex"].(int); ok {
		entity.MapIndex = value
	}
	if value, ok := object["x"].(int); ok {
		entity.X = value
	}
	if value, ok := object["y"].(int); ok {
		entity.Y = value
	}
}
