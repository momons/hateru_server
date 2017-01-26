package api

import (
	"../../constants"
	"../../entity/response"
	"../../manager"
	"../../util"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/satori/go.uuid"
	"net/http"
)

// トークン取得サービス.
type TokenGet struct {
	// アクセストークンマネージャ.
	accessTokensManager *manager.AccessTokens
	// ユーザマネージャ.
	usersManager *manager.Users
}

// インスタンス.
var instanceTokenGet *TokenGet

// インスタンス取得.
func GetTokenGet() *TokenGet {
	if instanceTokenGet == nil {
		instanceTokenGet = &TokenGet{
			accessTokensManager: manager.GetAccessToken(),
			usersManager:        manager.GetUsers(),
		}
	}
	return instanceTokenGet
}

// 受信.
func (service *TokenGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return
	}

	// ユーザコードチェック
	if !service.usersManager.HasUserCode(requestEntity.Status.UserCode) {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0002))
		return
	}

	// アクセストークン生成.
	accessToken := service.createAccessToken(requestEntity.Status.UserCode)

	// アクセストークンテーブルに追加
	isSuccess = service.accessTokensManager.DeleteInsert(requestEntity.Status.UserCode, accessToken)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4002))
		return
	}

	// レスポンス作成
	metaEntity := response.TokenGet{
		AccessToken: accessToken,
	}
	w.WriteJson(util.OKResponseEntity(metaEntity))

}

// アクセストークンを生成する.
func (service *TokenGet) createAccessToken(userCode string) string {

	accessToken := ""
	for {
		// アクセストークンを生成
		accessToken = uuid.NewV4().String()

		// 既存にあるかをチェック
		if !service.accessTokensManager.HasAccessToken(userCode, accessToken) {
			break
		}
	}

	return accessToken
}
