package manager

import (
	"../entity/database"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// 掲示板マネージャ.
type Bbs struct {
}

// インスタンス.
var instanceBbs *Bbs

// インスタンス取得.
func GetBbs() *Bbs {
	if instanceBbs == nil {
		instanceBbs = &Bbs{}
	}
	return instanceBbs
}

// 掲示板コードで取得
func (manager *Bbs) SelectList(
	bbsCode string,
	offset int,
	count int,
) *[]database.Bbs {

	var entities []database.Bbs

	err := Db.Where(
		"bbs_code = ?",
		bbsCode,
	).Order(
		"update_at DESC",
	).Offset(
		offset,
	).Limit(
		count,
	).Find(
		&entities,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// 掲示板コードとユーザコードとメッセージコードで取得
func (manager *Bbs) SelectOne(
	bbsCode string,
	userCode string,
	messageCode string,
) *database.Bbs {

	var entity database.Bbs

	err := Db.Where(
		"bbs_code = ? AND user_code = ? AND message_code = ?",
		bbsCode,
		userCode,
		messageCode,
	).First(
		&entity,
	).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

func (manager *Bbs) Insert(
	bbsCode string,
	userCode string,
	userName string,
	messageType string,
	messageData string,
) (bool, *string) {

	Db.Begin()

	// 現在日付取得
	nowAt := time.Now()

	// メッセージコード取得
	messageCode := uuid.NewV4().String()

	insertEntity := database.Bbs{
		BbsCode:     bbsCode,
		UserCode:    userCode,
		UserName:    userName,
		MessageCode: messageCode,
		MessageType: messageType,
		MessageData: messageData,
		UpdateAt:    nowAt,
		CreateAt:    nowAt,
	}
	err := Db.Create(&insertEntity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false, nil
	}

	Db.Commit()

	return true, &messageCode
}

// 削除.
func (manager *Bbs) Delete(
	bbsCode string,
	userCode string,
	messageCode string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(bbsCode, userCode, messageCode)
	if entity != nil {
		// 終了
		Db.Delete(&entity)
		Db.Commit()
	} else {
		Db.Rollback()
	}

	return true
}
