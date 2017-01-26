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

// アイテム交換送信サービス.
type ItemSend struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemSend *ItemSend

// インスタンス取得.
func GetItemSend() *ItemSend {
	if instanceItemSend == nil {
		instanceItemSend = &ItemSend{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemSend
}

// 受信.
func (service *ItemSend) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemSend{}
	params.Convert(requestEntity.Params)

	// 追加
	isSuccess, exchangeToken := service.itemManager.Insert(
		requestEntity.Status.UserCode,
		requestEntity.Status.UserName,
		params.ItemKindIndex,
		params.ItemCode,
		params.HopeItemKindIndex,
		params.PartnerUserCode,
		params.Password,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4009))
		return
	}

	// レスポンス作成
	metaEntity := response.ItemSend{
		ExchangeToken: *exchangeToken,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
