package api

import (
	"../../entity/request"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// プロフィール取得サービス.
type ProfileGet struct {
	// プロフィールマネージャ.
	profileManager *manager.UserProfiles
}

// インスタンス.
var instanceProfileGet *ProfileGet

// インスタンス取得.
func GetProfileGet() *ProfileGet {
	if instanceProfileGet == nil {
		instanceProfileGet = &ProfileGet{
			profileManager: manager.GetUserProfiles(),
		}
	}
	return instanceProfileGet
}

// 受信.
func (service *ProfileGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckToken(w, req)
	if !isSuccess {
		return
	}
	params := request.ProfileGet{}
	params.Convert(requestEntity.Params)

	// プロフィール取得
	userProfileEntity := service.profileManager.SelectOne(params.UserCode)
	if userProfileEntity == nil {
		w.WriteJson(util.ErrorResponseEntity(http.StatusNoContent, ""))
		return
	}

	// レスポンス作成
	metaEntity := response.ProfileGet{
		UserCode:    userProfileEntity.UserCode,
		UserName:    userProfileEntity.UserName,
		ProfileData: userProfileEntity.ProfileData,
	}
	// レスポンス返却
	w.WriteJson(util.OKResponseEntity(metaEntity))
}
