package response

import (
	"../database"
)

// 掲示板取得レスポンスEntity.
type BbsGet struct {
	// 掲示板ユーザコード.
	BbsCode string `json:"bbsCode"`
	// メッセージ.
	Messages []BbsGetMessage `json:"messages"`
}

type BbsGetMessage struct {
	// ユーザコード.
	UserCode string `json:"userCode"`
	// ユーザ名.
	UserName string `json:"userName"`
	// メッセージコード.
	MessageCode string `json:"messageCode"`
	// メッセージタイプ.
	MessageType string `json:"messageType"`
	// メッセージデータ.
	MessageData string `json:"messageData"`
	// 追加日時.
	CreateAt string `json:"createAt"`
}

// 掲示板データベースから変換
func (entity *BbsGetMessage) Convert(
	bbsEntity *database.Bbs,
	createAt string,
) {
	entity.UserCode = bbsEntity.UserCode
	entity.UserName = bbsEntity.UserName
	entity.MessageCode = bbsEntity.MessageCode
	entity.MessageType = bbsEntity.MessageType
	entity.MessageData = bbsEntity.MessageData
	entity.CreateAt = createAt
}
