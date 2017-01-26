package request

// プロフィール送信リクエストEntity.
type ProfileSend struct {
	// プロフィールデータ.
	ProfileData string `json:"profileData"`
}

func (entity *ProfileSend) Convert(
	object map[string]interface{},
) {
	if value, ok := object["profileData"].(string); ok {
		entity.ProfileData = value
	}
}
