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

// アイテム交換取得サービス.
type ItemGet struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemGet *ItemGet

// インスタンス取得.
func GetItemGet() *ItemGet {
	if instanceItemGet == nil {
		instanceItemGet = &ItemGet{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemGet
}

// 受信.
func (service *ItemGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemGet{}
	params.Convert(requestEntity.Params)

	// アイテム取得
	entity := service.itemManager.SelectOne(params.UserCode, params.ExchangeToken)
	if entity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}
	// パスワードチェック
	if !service.IsValidPassword(params, entity.PasswordHash) {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// 更新
	isSuccess = service.itemManager.UpdateExchangeStatus(
		params.UserCode,
		params.ExchangeToken,
		constants.ExchangeStatusTypeReplaced,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4008))
		return
	}

	// レスポンス作成
	metaEntity := response.ItemGet{
		ItemKindIndex: entity.ItemKindIndex,
		ItemCode:      entity.ItemCode,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}

// パスワードチェック
func (service *ItemGet) IsValidPassword(entity request.ItemGet, passwordHash string) bool {
	if len(entity.Password) <= 0 && len(passwordHash) <= 0 {
		// どちらもパスワードなし
		return true
	}
	if (len(entity.Password) <= 0 && len(passwordHash) > 0) || (len(entity.Password) > 0 && len(passwordHash) <= 0) {
		// どちらかが入っている
		return false
	}

	// パスワードをハッシュ化
	hashStr := util.Hash(entity.Password, entity.UserCode, entity.ExchangeToken)
	return hashStr == passwordHash
}
