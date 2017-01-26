package response

// セーブデータ取得レスポンスEntity.
type SaveGet struct {
	// セーブデータ.
	SaveData string `json:"saveData"`
	// チェックデジット.
	CheckDigit string `json:"checkDigit"`
}
