package request

// プロフィール取得リクエストEntity.
type ProfileGet struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
}

func (entity *ProfileGet) Convert(
	object map[string]interface{},
) {
	if value, ok := object["userCode"].(string); ok {
		entity.UserCode = value
	}
}
