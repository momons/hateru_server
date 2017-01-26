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

// アイテム交換リスト(自分向け)取得サービス.
type ItemListGetMy struct {
	// アイテムマネージャ.
	itemManager *manager.ExchangeItems
}

// インスタンス.
var instanceItemListGetMy *ItemListGetMy

// インスタンス取得.
func GetItemListGetMy() *ItemListGetMy {
	if instanceItemListGetMy == nil {
		instanceItemListGetMy = &ItemListGetMy{
			itemManager: manager.GetExchangeItems(),
		}
	}
	return instanceItemListGetMy
}

// 受信.
func (service *ItemListGetMy) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ItemListGetMy{}
	params.Convert(requestEntity.Params)

	// 自分向けの情報を取得
	entites := service.itemManager.SelectTargetMy(requestEntity.Status.UserCode)
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
	metaEntity := response.ItemListGetMy{
		Items: metaDetailEntity,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
