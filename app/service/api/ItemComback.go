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

// アイテム交換戻しサービス.
type ItemComback struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemComback *ItemComback

// インスタンス取得.
func GetItemComback() *ItemComback {
	if instanceItemComback == nil {
		instanceItemComback = &ItemComback{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemComback
}

// 受信.
func (service *ItemComback) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemComback{}
	params.Convert(requestEntity.Params)

	// 自分のアイテムを取得
	itemEntity := service.itemManager.SelectOne(
		requestEntity.Status.UserCode,
		params.ExchangeToken,
	)
	if itemEntity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// 状態"未交換"をチェック
	if itemEntity.ExchangeStatus != constants.ExchangeStatusTypeUnexchanged {
		// 未交換以外はダメ
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// 状態を"返却済み"に更新
	// TODO: 削除しようか考え中...　でも、弱電界でエラーになったら対応できないから消さないほうがいいかな？
	isSuccess = service.itemManager.UpdateExchangeStatus(
		requestEntity.Status.UserCode,
		params.ExchangeToken,
		constants.ExchangeStatusTypeReturned,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4008))
		return
	}

	// レスポンス作成
	metaEntity := response.ItemComback{
		ItemKindIndex: itemEntity.ItemKindIndex,
		ItemCode:      itemEntity.ItemCode,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
