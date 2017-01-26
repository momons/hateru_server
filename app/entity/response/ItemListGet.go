package response

// アイテムリスト取得レスポンスEntity.
type ItemListGet struct {
	// アイテム情報
	Items []ItemInfoGet `json:"items"`
}
