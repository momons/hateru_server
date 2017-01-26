package manager

import (
	"../entity/database"
	"log"
	"time"
)

// ユーザマネージャ.
type Users struct {
}

// インスタンス.
var instanceUsers *Users

// インスタンス取得.
func GetUsers() *Users {
	if instanceUsers == nil {
		instanceUsers = &Users{}
	}
	return instanceUsers
}

// ユーザコードで取得.
func (manager *Users) SelectOne(
	userCode string,
) *database.Users {

	var entity database.Users

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

// ユーザ既存チェック.
func (manager *Users) HasUserCode(
	userCode string,
) bool {
	entity := manager.SelectOne(userCode)
	return entity != nil
}

// 作成.
func (manager *Users) Insert(
	userCode string,
) bool {

	Db.Begin()

	nowAt := time.Now()

	entity := database.Users{
		UserCode: userCode,
		UpdateAt: nowAt,
		CreateAt: nowAt,
	}
	err := Db.Create(&entity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false
	}

	Db.Commit()

	return true
}
