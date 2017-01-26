package manager

import (
	"../constants"
	"../entity/database"
	"log"
	"time"
)

type AccessTokens struct {
}

// インスタンス.
var instanceAccessTokens *AccessTokens

// インスタンス取得.
func GetAccessToken() *AccessTokens {
	if instanceAccessTokens == nil {
		instanceAccessTokens = &AccessTokens{}
	}
	return instanceAccessTokens
}

// ユーザコード、トークンで取得
func (manager *AccessTokens) SelectOne(
	userCode string,
	accessToken string,
) *database.AccessTokens {

	var entity database.AccessTokens

	err := Db.Where(
		"user_code = ? AND access_token = ?",
		userCode,
		accessToken,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

func (manager *AccessTokens) HasAccessToken(
	userCode string,
	accessToken string,
) bool {
	entity := manager.SelectOne(userCode, accessToken)
	if entity == nil {
		return false
	}
	// 期限チェック
	if time.Now().Unix() > entity.PeriodAt.Unix() {
		return false
	}
	return true
}

// 作成.
func (manager *AccessTokens) DeleteInsert(
	userCode string,
	accessToken string,
) bool {

	Db.Begin()

	nowAt := time.Now()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode, accessToken)
	if entity != nil {
		Db.Delete(&entity)
	}

	insertEntity := database.AccessTokens{
		UserCode:    userCode,
		AccessToken: accessToken,
		PeriodAt:    nowAt.AddDate(0, 0, constants.AccessTokenPeriod),
		UpdateAt:    nowAt,
		CreateAt:    nowAt,
	}
	err := Db.Create(&insertEntity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false
	}

	Db.Commit()

	return true
}
