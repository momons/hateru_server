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

// 位置情報取得サービス.
type LocationGet struct {
	// 位置マネージャ.
	locationsManager *manager.Locations
}

// インスタンス.
var instanceLocationGet *LocationGet

// インスタンス取得.
func GetLocationGet() *LocationGet {
	if instanceLocationGet == nil {
		instanceLocationGet = &LocationGet{
			locationsManager: manager.GetLocations(),
		}
	}
	return instanceLocationGet
}

// 受信.
func (service *LocationGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.LocationGet{}
	params.Convert(requestEntity.Params)

	// 位置情報取得
	locationsEntities := service.locationsManager.SelectList(
		params.MapIndex,
		params.X,
		params.Y,
	)
	if locationsEntities == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4004))
		return
	}

	// レスポンス作成
	metaEntity := response.LocationGet{
		Locations: make([]response.LocationGetDetail, len(*locationsEntities)),
	}
	for index, entity := range *locationsEntities {
		metaEntity.Locations[index] = response.LocationGetDetail{
			UserCode:   entity.UserCode,
			UserName:   entity.UserName,
			Message:    entity.Message,
			MapIndex:   entity.MapIndex,
			X:          entity.X,
			Y:          entity.Y,
			OtherInfos: entity.OtherInfos,
		}
	}

	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
