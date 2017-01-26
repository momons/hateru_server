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

// プロフィール送信サービス.
type ProfileSend struct {
	// プロフィールマネージャ.
	profileManager *manager.UserProfiles
}

// インスタンス.
var instanceProfileSend *ProfileSend

// インスタンス取得.
func GetProfileSend() *ProfileSend {
	if instanceProfileSend == nil {
		instanceProfileSend = &ProfileSend{
			profileManager: manager.GetUserProfiles(),
		}
	}
	return instanceProfileSend
}

// 受信.
func (service *ProfileSend) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ProfileSend{}
	params.Convert(requestEntity.Params)

	// プロフィール追加
	isSuccess = service.profileManager.DeleteInsert(
		requestEntity.Status.UserCode,
		requestEntity.Status.UserName,
		params.ProfileData,
	)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4006))
		return
	}

	// レスポンス作成
	metaEntity := response.ProfileSend{}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
