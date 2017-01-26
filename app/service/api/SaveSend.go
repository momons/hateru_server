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

// 保存情報送信サービス.
type SaveSend struct {
	// セーブデータマネージャ.
	saveDataManager *manager.SaveDatas
}

// インスタンス.
var instanceSaveSend *SaveSend

// インスタンス取得.
func GetSaveSend() *SaveSend {
	if instanceSaveSend == nil {
		instanceSaveSend = &SaveSend{
			saveDataManager: manager.GetSaveDatas(),
		}
	}
	return instanceSaveSend
}

// 受信.
func (service *SaveSend) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.SaveSend{}
	params.Convert(requestEntity.Params)

	// セーブデータ追加
	var saveToken *string
	isSuccess, saveToken = service.saveDataManager.DeleteInsert(
		requestEntity.Status.UserCode,
		params.SaveData,
		params.CheckDigit,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4007))
		return
	}

	// レスポンス作成
	metaEntity := response.SaveSend{
		SaveToken: *saveToken,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
