package service

import (
	"chs/models"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

type checkinServiceImpl struct{}

func (h checkinServiceImpl) GetById(ctx context.Context, activityPb *Checkin) (*Checkin, error) {
	checkinModel := new(models.CheckinModel)
	copier.Copy(activityPb, checkinModel)
	insertId, err := checkinModel.Insert(checkinModel)
	activityPb.Id = insertId
	return activityPb, err
}

func (h checkinServiceImpl) List(ctx context.Context, in *Query) (*Checkinlist, error) {
	return nil, nil
}

func (h checkinServiceImpl) Insert(ctx context.Context, activityPb *Checkin) (*Checkin, error) {
	return nil, nil
}
