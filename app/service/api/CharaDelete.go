package api

import (
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
)

// キャラ削除サービス.
type CharaDelete struct {
	// キャラクタマネージャ.
	charactersManager *manager.Characters
}

// インスタンス.
var instanceCharaDelete *CharaDelete

// インスタンス取得.
func GetCharaDelete() *CharaDelete {
	if instanceCharaDelete == nil {
		instanceCharaDelete = &CharaDelete{
			charactersManager: manager.GetCharacters(),
		}
	}
	return instanceCharaDelete
}

// 受信.
func (service *CharaDelete) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.CharaDelete{}
	params.Convert(requestEntity.Params)

	// キャラ削除
	service.charactersManager.Delete(requestEntity.Status.UserCode)

	// レスポンス作成
	metaEntity := response.CharaDelete{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
