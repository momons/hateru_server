package response

// トークン取得レスポンスEntity.
type TokenGet struct {
	// アクセストークン.
	AccessToken string `json:"accessToken"`
}
