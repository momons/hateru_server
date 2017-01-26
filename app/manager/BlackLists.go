package manager

import (
	"../entity/database"
	"log"
	"time"
)

// ブラックリストマネージャ.
type BlackLists struct {
}

// インスタンス.
var instanceBlackLists *BlackLists

// インスタンス取得.
func GetBlackLists() *BlackLists {
	if instanceBlackLists == nil {
		instanceBlackLists = &BlackLists{}
	}
	return instanceBlackLists
}

// ユーザコードで取得
func (manager *BlackLists) SelectList(
	userCode string,
) *[]database.BlackLists {

	var entities []database.BlackLists

	err := Db.Where(
		"user_code = ?",
		userCode,
	).Order(
		"update_at DESC",
	).Find(
		&entities,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// ユーザコードとブラックユーザコードで取得.
func (manager *BlackLists) SelectOne(
	userCode string,
	blackUserCode string,
) *database.BlackLists {

	var entity database.BlackLists

	err := Db.Where(
		"user_code = ? AND black_user_code = ?",
		userCode,
		blackUserCode,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// ブラックリスト存在可否.
func (manager *BlackLists) HasBlackList(
	userCode string,
	blackUserCode string,
) bool {
	return manager.SelectOne(userCode, blackUserCode) != nil
}

// 作成.
func (manager *BlackLists) InsertIfNotExist(
	userCode string,
	blackUserCode string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode, blackUserCode)
	if entity != nil {
		// 終了
		Db.Rollback()
		return true
	}

	// 現在日付取得
	nowAt := time.Now()

	insertEntity := database.BlackLists{
		UserCode:      userCode,
		BlackUserCode: blackUserCode,
		UpdateAt:      nowAt,
		CreateAt:      nowAt,
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
func (manager *BlackLists) Delete(
	userCode string,
	blackUserCode string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode, blackUserCode)
	if entity != nil {
		// 終了
		Db.Delete(&entity)
		return true
	}

	Db.Commit()

	return true
}
