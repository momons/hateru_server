package manager

import (
	"../entity/database"
	"log"
	"time"
)

type Characters struct {
}

// インスタンス.
var instanceCharacters *Characters

// インスタンス取得.
func GetCharacters() *Characters {
	if instanceCharacters == nil {
		instanceCharacters = &Characters{}
	}
	return instanceCharacters
}

// ユーザコードで取得.
func (manager *Characters) SelectOne(
	userCode string,
) *database.Characters {

	var entity database.Characters

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

// 更新＆作成.
func (manager *Characters) InsertUpdate(
	userCode string,
	userName string,
	statusData string,
) bool {

	Db.Begin()

	// 現在日付取得
	nowAt := time.Now()

	// 既存チェック
	entity := manager.SelectOne(userCode)
	if entity != nil {
		entity.UserName = userName
		entity.StatusData = statusData
		entity.UpdateAt = nowAt
		err := Db.Update(&entity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false
		}
	} else {
		insertEntity := database.Characters{
			UserCode:   userCode,
			UserName:   userName,
			StatusData: statusData,
			UpdateAt:   nowAt,
			CreateAt:   nowAt,
		}
		err := Db.Create(&insertEntity).Error
		if err != nil {
			log.Println(err)
			Db.Rollback()
			return false
		}
	}

	Db.Commit()

	return true
}

// 削除.
func (manager *Characters) Delete(
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
