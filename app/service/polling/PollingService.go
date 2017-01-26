package polling

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

// ポーリングサービス.
type PollingService struct {
	// API
	api *rest.Api
}

// インスタンス.
var instancePollingService *PollingService

// インスタンス取得.
func GetPollingService() *PollingService {
	if instancePollingService == nil {
		instancePollingService = &PollingService{}
		instancePollingService.setup()
	}
	return instancePollingService
}

// ポーリングサービスを設定.
func (service *PollingService) setup() bool {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		// ポーリング
		&rest.Route{"POST", "/polling", service.recive},
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	api.SetApp(router)

	return true
}

func (service *PollingService) Start(port int) bool {

	// ポート番号を文字列化
	portStr := fmt.Sprintf("%d", port)

	log.Printf("Polling server started. port=%s", portStr)
	log.Fatal(http.ListenAndServe(":"+portStr, service.api.MakeHandler()))

	return true
}

func (service *PollingService) recive(w rest.ResponseWriter, req *rest.Request) {

}
