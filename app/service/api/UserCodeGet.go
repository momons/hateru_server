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

// ユーザコード取得サービス.
type UserCodeGet struct {
	// ユーザマネージャ.
	usersManager *manager.Users
}

// インスタンス.
var instanceUserCodeGet *UserCodeGet

// インスタンス取得.
func GetUserCodeGet() *UserCodeGet {
	if instanceUserCodeGet == nil {
		instanceUserCodeGet = &UserCodeGet{
			// ユーザマネージャ.
			usersManager: manager.GetUsers(),
		}
	}
	return instanceUserCodeGet
}

// 受信.
func (service *UserCodeGet) Recive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	isSuccess, _ := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return
	}

	// ユーザコード作成
	userCode := service.createUserCode()

	// ユーザテーブルに追加
	isSuccess = service.usersManager.Insert(userCode)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusInternalServerError, constants.MessageE4001))
		return
	}

	// レスポンス作成
	metaEntity := response.UserCodeGet{
		UserCode: userCode,
	}
	w.WriteJson(util.OKResponseEntity(metaEntity))
}

// ユーザコードを生成する.
func (service *UserCodeGet) createUserCode() string {

	userCode := ""
	for {
		// ユーザコード作成
		userCode = uuid.NewV4().String()

		// 一応既存にあるかをチェック
		if !service.usersManager.HasUserCode(userCode) {
			break
		}
	}

	return userCode
}
