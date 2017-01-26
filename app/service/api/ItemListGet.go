package api

import (
	_ "../../constants"
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	_ "net/http"
)

// アイテム交換リスト取得サービス.
type ItemListGet struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemListGet *ItemListGet

// インスタンス取得.
func GetItemListGet() *ItemListGet {
	if instanceItemListGet == nil {
		instanceItemListGet = &ItemListGet{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemListGet
}

// 受信.
func (service *ItemListGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemListGet{}
	params.Convert(requestEntity.Params)

	// リスト取得
	entites := service.itemManager.SelectSearch(params)
	metaDetailEntity := make([]response.ItemInfoGet, len(*entites))
	for index, entity := range *entites {
		metaDetailEntity[index] = response.ItemInfoGet{
			ItemKindIndex:     entity.ItemKindIndex,
			ItemCode:          entity.ItemCode,
			HopeItemKindIndex: entity.HopeItemKindIndex,
			ExchangeStatus:    entity.ExchangeStatus,
			HasPassword:       entity.HasPassword(),
		}
	}

	// レスポンス作成
	metaEntity := response.ItemListGet{
		Items: metaDetailEntity,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
