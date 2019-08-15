package dao

import (
	"chs/common"
	"chs/config"
	"time"
)

type CheckinModel struct {
	Id          int64 `xorm:"pk"`
	Wid         int64
	ActivityId  int64
	Wuid        int64
	Liner       int64
	Total       int64
	Lastcheckin time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (this *CheckinModel) TableName() string {
	return "checkins"
}

func (this *CheckinModel) ListByWid(wid int64) (lotteries []*CheckinModel) {
	if wid == 0 {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Find(&lotteries)
	if err != nil {
		return nil
	}
	return lotteries
}

/**
 * @GetCheckinByActivityWuid
 * @Description 活动用户签到信息
 * @Param id ActivityId
 * @Param id Wuid
 * @return CheckinModel,error
 */
func (this *CheckinModel) GetCheckinInfoByActivityIdAndWuid(activityId, wuid int64) (*CheckinModel, error) {
	if activityId == 0 || wuid == 0 {
		return nil, common.ErrDataGet
	}
	checkin := new(CheckinModel)
	checkin.ActivityId = activityId
	checkin.Wuid = wuid
	has, err := config.GetDbR(APP_DB_READ).Get(checkin)
	if !has || err != nil {
		return nil, err
	}
	return checkin, nil
}

/**
 * @GetCheckinByActivityWuid
 * @Description 活动用户签到信息，不存在自动创建
 * @Param id ActivityId
 * @Param id Wuid
 * @Param id Wid
 * @return CheckinModel,error
 */
func (this *CheckinModel) GetCheckinByActivityWuid(activityId, wuid int64) (*CheckinModel, error) {
	if activityId == 0 || wuid == 0 {
		return nil, common.ErrDataGet
	}
	checkin := new(CheckinModel)
	checkin.ActivityId = activityId
	checkin.Wuid = wuid
	has, err := config.GetDbR(APP_DB_READ).Get(checkin)
	if err != nil {
		return nil, common.ErrDataGet
	} else if !has {
		_, err = this.Insert(checkin)
		if err != nil {
			return nil, common.ErrDataGet
		}
	}
	return checkin, nil
}

func (this *CheckinModel) Insert(checkin *CheckinModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(checkin)
}

func (this *CheckinModel) Update(checkin *CheckinModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(checkin.Id).Update(checkin)
}

func (this *CheckinModel) DeleteById(id int64) bool {
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&CheckinModel{})
	if err != nil {
		return false
	}
	return true
}
