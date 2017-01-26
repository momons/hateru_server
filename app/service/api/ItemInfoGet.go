package api

import (
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// アイテム交換情報取得サービス.
type ItemInfoGet struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemInfoGet *ItemInfoGet

// インスタンス取得.
func GetItemInfoGet() *ItemInfoGet {
	if instanceItemInfoGet == nil {
		instanceItemInfoGet = &ItemInfoGet{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemInfoGet
}

// 受信.
func (service *ItemInfoGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemInfoGet{}
	params.Convert(requestEntity.Params)

	// アイテム情報取得
	entity := service.itemManager.SelectOne(params.UserCode, params.ExchangeToken)
	if entity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}
	// パートナーありの場合は検索できない
	if entity.IsHavePartner() {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// レスポンス作成
	metaEntity := response.ItemInfoGet{
		ItemKindIndex:     entity.ItemKindIndex,
		ItemCode:          entity.ItemCode,
		HopeItemKindIndex: entity.HopeItemKindIndex,
		ExchangeStatus:    entity.ExchangeStatus,
		HasPassword:       entity.HasPassword(),
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
