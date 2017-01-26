package request

// 掲示板取得リクエストEntity.
type BbsGet struct {
	// 掲示板コード.
	BbsCode string `json:"bbsCode"`
	// 取得オフセット.
	Offset int `json:"offset"`
	// 取得カウント.
	Count int `json:"count"`
}

// コンバート.
func (entity *BbsGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["bbsCode"].(string); ok {
		entity.BbsCode = value
	}
	if value, ok := object["offset"].(int); ok {
		entity.Offset = value
	}
	if value, ok := object["count"].(int); ok {
		entity.Count = value
	}
}
