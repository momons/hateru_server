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

// キャラ取得サービス.
type CharaGet struct {
	// キャラクタマネージャ.
	charactersManager *manager.Characters
}

// インスタンス.
var instanceCharaGet *CharaGet

// インスタンス取得.
func GetCharaGet() *CharaGet {
	if instanceCharaGet == nil {
		instanceCharaGet = &CharaGet{
			charactersManager: manager.GetCharacters(),
		}
	}
	return instanceCharaGet
}

// 受信.
func (service *CharaGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.CharaGet{}
	params.Convert(requestEntity.Params)

	// 取得
	charactersEntity := service.charactersManager.SelectOne(params.UserCode)
	if charactersEntity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, constants.MessageE2001))
		return
	}

	// レスポンス作成
	metaEntity := response.CharaGet{
		UserCode:   charactersEntity.UserCode,
		UserName:   charactersEntity.UserName,
		StatusData: charactersEntity.StatusData,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
