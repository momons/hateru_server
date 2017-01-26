package response

// アイテム取り戻しレスポンスEntity.
type ItemComback struct {
	// アイテム種類.
	ItemKindIndex int `json:"itemKindIndex"`
	// アイテムコード.
	ItemCode string `json:"itemCode"`
}
