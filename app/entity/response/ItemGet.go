package response

// アイテム取得レスポンスEntity.
type ItemGet struct {
	// アイテム種類.
	ItemKindIndex int `json:"itemKindIndex"`
	// 所持アイテムコード.
	ItemCode string `json:"itemCode"`
}
