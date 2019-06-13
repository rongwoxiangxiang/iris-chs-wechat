package models

import (
	"iris/common"
	"time"
)

type PrizeHistoryModel struct {
	Id         int64 `xorm:"pk"`
	ActivityId int64
	Wuid       int64
	Prize      string
	Code       string
	Level      int8
	CreatedAt  time.Time
}

func (this *PrizeHistoryModel) TableName() string {
	return "prize_history"
}

func (this *PrizeHistoryModel) GetByActivityWuId(activityId, wuid int64) (*PrizeHistoryModel, error) {

	history := new(PrizeHistoryModel)
	history.Wuid = wuid
	history.ActivityId = activityId
	has, err := Db.Desc("id").Get(history)
	if err != nil {
		return nil, common.ErrDataGet
	} else if has == false {
		return nil, common.ErrDataEmpty
	}
	return history, nil
}

func (this *PrizeHistoryModel) LimitUnderActivityList(activityId int64, index int, limit int) (histories []*PrizeHistoryModel) {
	if activityId == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := Db.Where("acitivity_id = ?", activityId).Limit(limit, (index-1)*limit).Find(&histories)
	if err != nil {
		return nil
	}
	return histories
}

func (this *PrizeHistoryModel) Insert(model *PrizeHistoryModel) (int64, error) {
	return Db.InsertOne(model)
}

func (this *PrizeHistoryModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := Db.Id(id).Unscoped().Delete(&PrizeHistoryModel{})
	if err != nil {
		return false
	}
	return true
}
