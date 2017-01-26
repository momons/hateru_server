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

// 位置情報送信サービス.
type LocationSend struct {
	// 位置マネージャ.
	locationsManager *manager.Locations
}

// インスタンス.
var instanceLocationSend *LocationSend

// インスタンス取得.
func GetLocationSend() *LocationSend {
	if instanceLocationSend == nil {
		instanceLocationSend = &LocationSend{
			locationsManager: manager.GetLocations(),
		}
	}
	return instanceLocationSend
}

// 受信.
func (service *LocationSend) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.LocationSend{}
	params.Convert(requestEntity.Params)

	// 追加
	isSuccess = service.locationsManager.DeleteInsert(
		requestEntity.Status.UserCode,
		requestEntity.Status.UserName,
		params,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4005))
		return
	}

	// レスポンス作成
	metaEntity := response.LocationSend{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
