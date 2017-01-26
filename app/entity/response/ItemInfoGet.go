package response

// アイテム情報取得レスポンスEntity.
type ItemInfoGet struct {
	// アイテム種類.
	ItemKindIndex int `json:"itemKindIndex"`
	// アイテム種類.
	ItemCode string `json:"itemCode"`
	// 交換希望アイテム種別.
	HopeItemKindIndex int `json:"hopeItemKindIndex"`
	// 交換ステータス.
	ExchangeStatus string `json:"exchangeStatus"`
	// パスワードあり
	HasPassword bool `json:"hasPassowrd"`
}
