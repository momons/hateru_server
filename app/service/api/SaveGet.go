package api

import (
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// 保存情報取得サービス.
type SaveGet struct {
	// セーブデータマネージャ.
	saveDataManager *manager.SaveDatas
}

// インスタンス.
var instanceSaveGet *SaveGet

// インスタンス取得.
func GetSaveGet() *SaveGet {
	if instanceSaveGet == nil {
		instanceSaveGet = &SaveGet{
			saveDataManager: manager.GetSaveDatas(),
		}
	}
	return instanceSaveGet
}

// 受信.
func (service *SaveGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.SaveGet{}
	params.Convert(requestEntity.Params)

	// セーブデータ取得
	saveDataEntity := service.saveDataManager.SelectUserCodeSaveToken(
		requestEntity.Status.UserCode,
		params.SaveToken,
	)
	if saveDataEntity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// レスポンス作成
	metaEntity := response.SaveGet{
		SaveData:   saveDataEntity.SaveData,
		CheckDigit: saveDataEntity.CheckDigit,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
