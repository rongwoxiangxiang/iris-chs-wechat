package models

import (
	"iris/common"
	"time"
)

type PrizeModel struct {
	Id         int64 `xorm:"pk"`
	Wid        int64
	ActivityId int64
	Code       string
	Level      int8
	Used       int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (this *PrizeModel) TableName() string {
	return "prizes"
}

func (this *PrizeModel) ChooseOneUsedPrize(activityId int64, level string, idGt int64) (prize *PrizeModel, err error) {
	if idGt > 0 {
		DbR.Where("id >= ?", idGt)
	}
	has, err := DbR.Where("activity_id = ? AND level = ? AND used = ?", activityId, level, common.NO_VALUE).Get(prize)
	if err != nil || has == false {
		return nil, common.ErrDataUnExist
	}
	prize.Used = common.YES_VALUE
	_, err = DbW.
		Where("id = ? and used = ?", prize.Id, common.NO_VALUE).
		Cols("used").
		Update(prize)
	if err != nil {
		return nil, common.ErrDataUpdate
	}
	return
}

func (this *PrizeModel) Insert(prize *PrizeModel) (int64, error) {
	return DbW.InsertOne(prize)
}

func (this *PrizeModel) InsertBatch(prizes []*PrizeModel) (int64, error) {
	return DbW.Insert(&prizes)
}

func (this *PrizeModel) DeleteById(id int64) bool {
	if id == 0 {
		return false
	}
	_, err := DbW.Id(id).Unscoped().Delete(&PrizeModel{})
	if err != nil {
		return false
	}
	return true
}

func (this *PrizeModel) LimitUnderActivityList(activityId int64, index int, limit int) (prizes []*PrizeModel) {
	if activityId == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := DbR.Where("acitivity_id = ?", activityId).Limit(limit, (index-1)*limit).Find(&prizes)
	if err != nil {
		return nil
	}
	return prizes
}
