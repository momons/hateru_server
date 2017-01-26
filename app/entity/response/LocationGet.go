package response

// 位置情報取得レスポンスEntity.
type LocationGet struct {
	// 位置情報群.
	Locations []LocationGetDetail `json:"locations"`
}

// 位置情報取得詳細レスポンスEntity.
type LocationGetDetail struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
	// ユーザ名.
	UserName string `json:"userName"`
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
