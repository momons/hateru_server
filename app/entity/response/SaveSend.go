package response

// セーブデータ送信レスポンスEntity.
type SaveSend struct {
	// セーブトークン.
	SaveToken string `json:"saveToken"`
}
