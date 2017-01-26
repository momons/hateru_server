package response

// 自分向けアイテムリスト取得レスポンスEntity.
type ItemListGetMy struct {
	// アイテム情報
	Items []ItemInfoGet `json:"items"`
}
