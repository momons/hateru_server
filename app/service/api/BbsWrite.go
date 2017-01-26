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

// 掲示板書き込みサービス.
type BbsWrite struct {
	// 掲示板マネージャ.
	bbsManager *manager.Bbs
	// ブラックリストマネージャ.
	blackListsManager *manager.BlackLists
}

// インスタンス.
var instanceBbsWrite *BbsWrite

// インスタンス取得.
func GetBbsWrite() *BbsWrite {
	if instanceBbsWrite == nil {
		instanceBbsWrite = &BbsWrite{
			bbsManager:        manager.GetBbs(),
			blackListsManager: manager.GetBlackLists(),
		}
	}
	return instanceBbsWrite
}

// 受信.
func (service *BbsWrite) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.BbsWrite{}
	params.Convert(requestEntity.Params)

	// ブラックリストチェック
	isExist := service.blackListsManager.HasBlackList(
		params.BbsCode,
		requestEntity.Status.UserCode,
	)
	if isExist {
		w.WriteJson(util.ErrorResponseEntity(http.StatusForbidden, constants.MessageE1001))
		return
	}

	//// 掲示板読み込み
	//bbsEntities := service.bbsManager.SelectList(
	//	params.BbsCode,
	//	params.Offset,
	//	params.Count,
	//)

	// レスポンス作成
	metaEntity := response.BbsWrite{}
	//if bbsEntities != nil {
	//	metaEntity.BbsCode = params.BbsCode
	//	metaEntity.Messages = make([]response.BbsGetMessage, len(*bbsEntities))
	//	for index, bbsEntity := range *bbsEntities {
	//		metaEntity.Messages[index].Convert(&bbsEntity)
	//	}
	//}

	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
