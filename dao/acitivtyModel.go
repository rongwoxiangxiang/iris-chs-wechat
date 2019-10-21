package dao

import (
	"chs/config"
	"strconv"
	"time"
)

type ActivityInterfaceR interface {
	GetById(int64) *ActivityModel
	LimitUnderWidList(index, limit, wid int) []*ActivityModel
}

type ActivityInterfaceW interface {
	Insert(*ActivityModel) (int64, error)
	Update(*ActivityModel) (int64, error)
	DeleteById(int64) bool
}

type ActivityModel struct {
	Id           int64 `xorm:"pk"`
	Wid          int64
	Name         string
	Desc         string
	ActivityType int8
	RelativeId   int64
	Extra        string
	TimeStarted  time.Time
	TimeEnd      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (this *ActivityModel) TableName() string {
	return "activities"
}

func (this *ActivityModel) GetById(id int64) *ActivityModel {
	if id != 0 {
		activity := new(ActivityModel)
		config.CacheGetStruct(this.TableName()+strconv.FormatInt(id, 10), activity)
		if activity != nil && activity.Id > 0 {
			return activity
		}
		activity.Id = id
		has, err := config.GetDbR(APP_DB_READ).Get(activity)
		if has != true || err != nil {
			config.Logger().Error("ActivityModel GetById [%d] empty or err:", id, err)
			return nil
		}
		config.CacheSetJson(this.TableName()+strconv.FormatInt(id, 10), activity, 3600*24*10)
		return activity
	}
	return nil
}

func (this *ActivityModel) LimitUnderWidList(index, limit, wid int) (activities []*ActivityModel) {
	if index < 1 || wid < 1 || limit < 1 {
		return nil
	}
	err := config.GetDbR(APP_DB_READ).Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&activities)
	if err != nil {
		config.Logger().Error("ActivityModel LimitUnderWidList err:", err)
		return nil
	}
	return activities
}

func (this *ActivityModel) Insert(activity *ActivityModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).InsertOne(activity)
}

func (this *ActivityModel) Update(activity *ActivityModel) (int64, error) {
	return config.GetDbW(APP_DB_WRITE).Id(activity.Id).Update(activity)
}

func (this *ActivityModel) DeleteById(id int64) bool {
	if id < 1 {
		return false
	}
	_, err := config.GetDbW(APP_DB_WRITE).Id(id).Unscoped().Delete(&ActivityModel{})
	if err != nil {
		config.Logger().Error("ActivityModel DeleteById err:", err)
		return false
	}
	return true
}
