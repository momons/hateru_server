package response

// キャラ取得レスポンスEntity.
type CharaGet struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
	// ユーザ名.
	UserName string `json:"userName"`
	// ステータスデータ.
	StatusData string `json:"statusData"`
}
