package constants

const (
	// トークンの期限 7日
	AccessTokenPeriod = 7

	// 交換ステータスタイプ 未交換
	ExchangeStatusTypeUnexchanged = "unexchanged"
	// 交換ステータスタイプ 交換済み
	ExchangeStatusTypeReplaced = "replaced"
	// 交換ステータスタイプ 返却済み
	ExchangeStatusTypeReturned = "returned"
)
