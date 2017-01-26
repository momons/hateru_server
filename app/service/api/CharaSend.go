package api

import (
	"../../constants"
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// キャラ送信サービス.
type CharaSend struct {
	// キャラクタマネージャ.
	charactersManager *manager.Characters
}

// インスタンス.
var instanceCharaSend *CharaSend

// インスタンス取得.
func GetCharaSend() *CharaSend {
	if instanceCharaSend == nil {
		instanceCharaSend = &CharaSend{}
	}
	return instanceCharaSend
}

// 受信.
func (service *CharaSend) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.CharaSend{}
	params.Convert(requestEntity.Params)

	// 追加
	isSuccess = service.charactersManager.InsertUpdate(
		requestEntity.Status.UserCode,
		requestEntity.Status.UserName,
		params.StatusData,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4003))
		return
	}

	// レスポンス作成
	metaEntity := response.CharaSend{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
