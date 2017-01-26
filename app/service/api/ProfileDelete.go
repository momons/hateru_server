package api

import (
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
)

// プロフィール削除サービス.
type ProfileDelete struct {
	// プロフィールマネージャ.
	profileManager *manager.UserProfiles
}

// インスタンス.
var instanceProfileDelete *ProfileDelete

// インスタンス取得.
func GetProfileDelete() *ProfileDelete {
	if instanceProfileDelete == nil {
		instanceProfileDelete = &ProfileDelete{
			profileManager: manager.GetUserProfiles(),
		}
	}
	return instanceProfileDelete
}

// 受信.
func (service *ProfileDelete) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ProfileDelete{}
	params.Convert(requestEntity.Params)

	// プロフィール削除
	service.profileManager.Delete(requestEntity.Status.UserCode)

	// レスポンス作成
	metaEntity := response.ProfileDelete{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
