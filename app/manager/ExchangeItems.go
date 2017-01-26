package manager

import (
	"../constants"
	"../entity/database"
	"../entity/request"
	"../util"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// アイテム交換マネージャ.
type ExchangeItems struct {
}

// インスタンス.
var instanceExchangeItems *ExchangeItems

// インスタンス取得.
func GetExchangeItems() *ExchangeItems {
	if instanceExchangeItems == nil {
		instanceExchangeItems = &ExchangeItems{}
	}
	return instanceExchangeItems
}

// ユーザコードで取得.
func (manager *ExchangeItems) SelectTargetMy(
	userCode string,
) *[]database.ExchangeItems {

	var entities []database.ExchangeItems

	err := Db.Where(
		"partner_user_code = ? AND exchange_status = ?",
		userCode,
		constants.ExchangeStatusTypeUnexchanged,
	).Order(
		"update_at desc",
	).Find(&entities).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// 条件で取得.
func (manager *ExchangeItems) SelectSearch(
	requestEntity request.ItemListGet,
) *[]database.ExchangeItems {

	var entities []database.ExchangeItems

	// WHERE句作成
	whereStr := "partner_user_code = ''"
	whereStr += " AND exchange_status = '" + constants.ExchangeStatusTypeUnexchanged + "'"
	if requestEntity.IsTransfer() {
		whereStr += " AND hope_item_kind_index < 0"
	} else if requestEntity.IsExchange() {
		whereStr += " AND hope_item_kind_index >= 0"
	}
	if requestEntity.ItemKindIndex >= 0 {
		whereStr += " AND item_kind_index = " + string(requestEntity.ItemKindIndex)
	}
	if requestEntity.HopeItemKindIndex >= 0 {
		whereStr += " AND hope_item_kind_index = " + string(requestEntity.HopeItemKindIndex)
	}

	err := Db.Where(
		whereStr,
	).Order(
		"update_at desc",
	).Offset(
		requestEntity.Offset,
	).Limit(
		requestEntity.Count,
	).Find(&entities).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// ユーザコード、交換トークンで取得.
func (manager *ExchangeItems) SelectOne(
	userCode string,
	exchangeToken string,
) *database.ExchangeItems {

	var entity database.ExchangeItems

	err := Db.Where(
		"user_code = ? AND exchange_token = ?",
		userCode,
		exchangeToken,
	).First(&entity).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entity
}

// 作成.
func (manager *ExchangeItems) Insert(
	userCode string,
	userName string,
	itemKindIndex int,
	itemCode string,
	hopeItemKindIndex int,
	partnerUserCode string,
	password string,
) (bool, *string) {

	Db.Begin()

	// 現在日付取得
	nowAt := time.Now()
	// セーブトークン取得
	exchangeToken := uuid.NewV4().String()
	// パスワードをハッシュ化
	passwordHash := util.Hash(password, userCode, exchangeToken)

	insertEntity := database.ExchangeItems{
		UserCode:          userCode,
		ExchangeToken:     exchangeToken,
		UserName:          userName,
		ItemKindIndex:     itemKindIndex,
		ItemCode:          itemCode,
		HopeItemKindIndex: hopeItemKindIndex,
		PartnerUserCode:   partnerUserCode,
		ExchangeStatus:    constants.ExchangeStatusTypeUnexchanged,
		PasswordHash:      passwordHash,
		UpdateAt:          nowAt,
		CreateAt:          nowAt,
	}
	err := Db.Create(&insertEntity).Error
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false, nil
	}

	Db.Commit()

	return true, &exchangeToken
}

// ステータス更新.
func (manager *ExchangeItems) UpdateExchangeStatus(
	userCode string,
	exchangeToken string,
	exchangeStatus string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode, exchangeToken)
	if entity == nil {
		Db.Rollback()
		return false
	}

	// 更新
	entity.ExchangeStatus = exchangeStatus
	entity.UpdateAt = time.Now()

	// 更新
	err := Db.Update(&entity)
	if err != nil {
		log.Println(err)
		Db.Rollback()
		return false
	}

	// コミット
	Db.Commit()

	return true
}

// 削除.
func (manager *ExchangeItems) Delete(
	userCode string,
	exchangeToken string,
) bool {

	Db.Begin()

	// 既存にある場合削除
	entity := manager.SelectOne(userCode, exchangeToken)
	if entity != nil {
		Db.Delete(&entity)
		Db.Commit()
	} else {
		Db.Rollback()
	}

	return true
}
