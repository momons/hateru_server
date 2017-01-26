package manager

import (
	"../entity/database"
	"../entity/request"
	"log"
	"time"
)

// 位置情報マネージャ.
type Locations struct {
}

// インスタンス.
var instanceLocations *Locations

// インスタンス取得.
func GetLocations() *Locations {
	if instanceLocations == nil {
		instanceLocations = &Locations{}
	}
	return instanceLocations
}

// 範囲内のユニットを取得
func (manager *Locations) SelectList(
	mapIndex int,
	x int,
	y int,
) *[]database.Locations {

	var entities []database.Locations

	err := Db.Where(
		"map_index = ? AND x >= ? AND x <= ? AND y >= ? AND y <= ?",
		mapIndex,
		x-10,
		x+10,
		y-10,
		y+10,
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

// ユーザコードで取得
func (manager *Locations) SelectOne(
	userCode string,
) *database.Locations {

	var entity database.Locations

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
func (manager *Locations) DeleteInsert(
	userCode string,
	userName string,
	requestEntity request.LocationSend,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode)
	if entity != nil {
		Db.Delete(&entity)
	}

	// 現在日付取得
	nowAt := time.Now()

	insertEntity := database.Locations{
		UserCode:   userCode,
		UserName:   userName,
		Message:    requestEntity.Message,
		MapIndex:   requestEntity.MapIndex,
		X:          requestEntity.X,
		Y:          requestEntity.Y,
		OtherInfos: requestEntity.OtherInfos,
		UpdateAt:   nowAt,
		CreateAt:   nowAt,
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
