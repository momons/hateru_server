package manager

import (
	"../constants"
	"../entity/database"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

// データベース.
var Db *gorm.DB

// データベース設定.
func SetupDatabase() bool {

	// 接続文字列
	connectStr := "sslmode=disable host=localhost" + " dbname=" + constants.DatabaseName + " user=" + constants.UserId + " password=" + constants.Password
	db, err := gorm.Open("postgres", connectStr)
	if err != nil {
		log.Println(err)
		return false
	}
	db.DB()

	// デバッグモード.
	//db.LogMode(getParamBool("db.debug"))

	// 退避
	Db = db

	// テーブル設定.
	database.SetupAccessTokens(db)
	database.SetupBbs(db)
	database.SetupBlackLists(db)
	database.SetupCharacters(db)
	database.SetupExchangeItems(db)
	database.SetupExhibitBanItems(db)
	database.SetupLocations(db)
	database.SetupSaveDatas(db)
	database.SetupUserProfiles(db)
	database.SetupUsers(db)

	return true
}
