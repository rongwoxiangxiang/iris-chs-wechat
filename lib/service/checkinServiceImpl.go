package service

import (
	"chs/models"
	"context"
	"github.com/jinzhu/copier"
)

type CheckinServiceImpl struct{}

func (h CheckinServiceImpl) GetById(ctx context.Context, checkin *Checkin) (*Checkin, error) {
	checkinModel := new(models.CheckinModel)
	copier.Copy(checkin, checkinModel)
	insertId, err := checkinModel.Insert(checkinModel)
	checkin.Id = insertId
	return checkin, err
}

func (h CheckinServiceImpl) List(ctx context.Context, in *CheckinQuery) (*Checkinlist, error) {
	return nil, nil
}

func (h CheckinServiceImpl) Insert(ctx context.Context, activityPb *Checkin) (*Checkin, error) {
	return nil, nil
}
