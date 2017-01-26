package manager

import (
	"../entity/database"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// セーブデータマネージャ.
type SaveDatas struct {
}

// インスタンス.
var instanceSaveDatas *SaveDatas

// インスタンス取得.
func GetSaveDatas() *SaveDatas {
	if instanceSaveDatas == nil {
		instanceSaveDatas = &SaveDatas{}
	}
	return instanceSaveDatas
}

// ユーザコード、セーブトークンで取得
func (manager *SaveDatas) SelectUserCodeSaveToken(
	userCode string,
	saveToken string,
) *database.SaveDatas {

	var entity database.SaveDatas

	err := Db.Where("user_code = ? AND save_token = ?",
		userCode,
		saveToken,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// ユーザコードで取得
func (manager *SaveDatas) SelectOne(
	userCode string,
) *database.SaveDatas {

	var entity database.SaveDatas

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
func (manager *SaveDatas) DeleteInsert(
	userCode string,
	saveData string,
	checkDigit string,
) (bool, *string) {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode)
	if entity != nil {
		Db.Delete(&entity)
	}

	// 現在日付取得
	nowAt := time.Now()
	// セーブトークン取得
	saveToken := uuid.NewV4().String()

	insertEntity := database.SaveDatas{
		UserCode:   userCode,
		SaveToken:  saveToken,
		SaveData:   saveData,
		CheckDigit: checkDigit,
		UpdateAt:   nowAt,
		CreateAt:   nowAt,
	}
	err := Db.Create(&insertEntity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false, nil
	}

	Db.Commit()

	return true, &saveToken
}
