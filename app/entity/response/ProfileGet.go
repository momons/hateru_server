package response

// プロフィール取得レスポンスEntity.
type ProfileGet struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
	// ユーザ名.
	UserName string `json:"userName"`
	// プロフィールデータ.
	ProfileData string `json:"profileData"`
}
