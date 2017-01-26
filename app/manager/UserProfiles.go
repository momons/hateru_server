package manager

import (
	"../entity/database"
	"log"
	"time"
)

// ユーザプロフィールマネージャ.
type UserProfiles struct {
}

// インスタンス.
var instanceUserProfiles *UserProfiles

// インスタンス取得.
func GetUserProfiles() *UserProfiles {
	if instanceUserProfiles == nil {
		instanceUserProfiles = &UserProfiles{}
	}
	return instanceUserProfiles
}

// ユーザコードで取得
func (manager *UserProfiles) SelectOne(
	userCode string,
) *database.UserProfiles {

	var entity database.UserProfiles

	err := Db.Where(
		"user_code = ?",
		userCode,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// 削除＆作成.
func (manager *UserProfiles) DeleteInsert(
	userCode string,
	userName string,
	profileData string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode)
	if entity != nil {
		Db.Delete(&entity)
	}

	// 現在日付取得
	nowAt := time.Now()

	insertEntity := database.UserProfiles{
		UserCode:    userCode,
		UserName:    userName,
		ProfileData: profileData,
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

// 削除.
func (manager *UserProfiles) Delete(
	userCode string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode)
	if entity != nil {
		Db.Delete(&entity)
		Db.Commit()
	} else {
		Db.Rollback()
	}

	return true
}
