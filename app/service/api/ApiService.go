package api

import (
	"../../constants"
	"../../entity/request"
	"../../manager"
	"../../util"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

// APIサービス.
type ApiService struct {
	// API
	api *rest.Api
}

// インスタンス.
var instanceApiService *ApiService

// インスタンス取得.
func GetApiService() *ApiService {
	if instanceApiService == nil {
		instanceApiService = &ApiService{}
		instanceApiService.setup()
	}
	return instanceApiService
}

// APIサービスを設定.
func (service *ApiService) setup() bool {

	service.api = rest.NewApi()
	service.api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// ユーザコード取得
		&rest.Route{"POST", "/user_code_get", GetUserCodeGet().Recive},
		// トークン取得
		&rest.Route{"POST", "/token_get", GetTokenGet().Recive},
		// アイテム交換
		&rest.Route{"POST", "/item_send", GetItemGet().Recive},
		&rest.Route{"POST", "/item_comeback", GetItemComback().Recive},
		&rest.Route{"POST", "/item_get", GetItemGet().Recive},
		&rest.Route{"POST", "/item_list_get", GetItemListGet().Recive},
		&rest.Route{"POST", "/item_list_get_my", GetItemListGetMy().Recive},
		&rest.Route{"POST", "/item_info_get", GetItemInfoGet().Recive},
		// 位置情報
		&rest.Route{"POST", "/location_send", GetLocationSend().Recive},
		&rest.Route{"POST", "/location_get", GetLocationSend().Recive},
		// 掲示板
		&rest.Route{"POST", "/bbs_get", GetBbsGet().Recive},
		&rest.Route{"POST", "/bbs_write", GetBbsWrite().Recive},
		// キャラ
		&rest.Route{"POST", "/chara_send", GetCharaSend().Recive},
		&rest.Route{"POST", "/chara_delete", GetCharaDelete().Recive},
		&rest.Route{"POST", "/chara_get", GetCharaGet().Recive},
		// セーブデータ
		&rest.Route{"POST", "/save_send", GetSaveSend().Recive},
		&rest.Route{"POST", "/save_get", GetSaveGet().Recive},
		// プロフィール
		&rest.Route{"POST", "/profile_send", GetProfileSend().Recive},
		&rest.Route{"POST", "/profile_delete", GetProfileDelete().Recive},
		&rest.Route{"POST", "/profile_get", GetProfileGet().Recive},
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	service.api.SetApp(router)

	return true
}

// スタート.
func (service *ApiService) Start(port int) {
	// ポート番号を文字列化
	portStr := fmt.Sprintf("%d", port)

	log.Printf("API server started. port=%s", portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, service.api.MakeHandler()))
}

// リクエスト取得＆トークンチェック
func GetRequestAndCheckToken(
	w rest.ResponseWriter,
	req *rest.Request,
) (bool, *request.Common) {

	// リクエスト取得
	isSuccess, requestEntity := GetRequestAndCheckAppToken(w, req)
	if !isSuccess {
		return false, requestEntity
	}

	// アクセストークンチェック.
	if !manager.GetAccessToken().HasAccessToken(requestEntity.Status.UserCode, requestEntity.Status.AccessToken) {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0004))
		return false, requestEntity
	}

	return true, requestEntity
}

// リクエスト取得＆アプリトークンチェック
func GetRequestAndCheckAppToken(
	w rest.ResponseWriter,
	req *rest.Request,
) (bool, *request.Common) {

	// リクエスト取得
	requestEntity := request.Common{}
	isSuccess := util.RequestEntity(&requestEntity, req)
	if !isSuccess {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0001))
		return false, nil
	}

	// アプリトークンチェック
	if !requestEntity.Status.IsValidAppToken() {
		w.WriteJson(util.ErrorResponseEntity(http.StatusBadRequest, constants.MessageE0003))
		return false, &requestEntity
	}

	return true, &requestEntity
}
