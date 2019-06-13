package models

import (
	"iris/common"
	"time"
)

type ReplyModel struct {
	Id            int64 `xorm:"pk"`
	Wid           int64
	ActivityId    int64
	Alias         string
	ClickKey      string
	Success       string
	Fail          string //活动数据不存在，未找到等报错是返回信息
	NoPrizeReturn string
	Extra         string
	Type          string
	Disabled      int8
	Match         int8
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (r *ReplyModel) TableName() string {
	return "replies"
}

func (r *ReplyModel) GetById(id int64) *ReplyModel {
	if id != 0 {
		user := new(ReplyModel)
		user.Id = id
		has, err := Db.Get(user)
		if !has || err != nil {
			return nil
		}
		return user
	}
	return nil
}

/**
 * @Find
 * @Param Reply.Id int
 * @Param Reply.Alias string
 * @Param Reply.ClickKey string
 * @Success Reply
 */
func (r *ReplyModel) FindOne(model *ReplyModel) (reply *ReplyModel) {
	if model.Wid == 0 || ("" == model.Alias && "" == model.ClickKey) {
		return
	}
	qs := Db.Where("wid = ?", model.Wid)
	if "" != model.Alias {
		qs = qs.Where("alias = ?", model.Alias)
	} else if "" != r.ClickKey {
		qs = qs.Where("click_key = ?", model.ClickKey)
	}
	has, err := qs.Where("disabled = ?", common.NO_VALUE).Get(reply)
	if !has || err != nil {
		return nil
	}
	return
}

func (r *ReplyModel) LimitUnderWidList(wid int64, index int, limit int) (relpies []*ReplyModel) {
	if wid == 0 || (index < 1 && limit < 1) {
		return nil
	}
	err := Db.Where("wid = ?", wid).Limit(limit, (index-1)*limit).Find(&relpies)
	if err != nil {
		return nil
	}
	return relpies
}

func (r *ReplyModel) ChangeDisabledByWidActivityId(wid, activityId int64, disabled int8) bool {
	if wid == 0 || activityId == 0 {
		return false
	}
	reply := ReplyModel{Wid: wid, ActivityId: activityId}
	has, err := Db.Get(&reply)
	if err != nil || has == false {
		return false
	}
	reply.Disabled = disabled
	_, err = Db.Id(reply.Id).Cols("disabled").Update(reply)
	if err != nil {
		return false
	}
	return true
}

func (r *ReplyModel) Insert(model *ReplyModel) (int64, error) {
	return Db.InsertOne(model)
}

func (r *ReplyModel) Update(model *ReplyModel) (int64, error) {
	return Db.Id(model.Id).Update(model)
}

func (r *ReplyModel) DeleteById(id int64) bool {
	_, err := Db.Id(id).Unscoped().Delete(new(ReplyModel))
	if err != nil {
		return false
	}
	return true
}
