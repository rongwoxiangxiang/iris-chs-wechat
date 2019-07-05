package service

import (
	"chs/models"
	"github.com/jinzhu/copier"
	"golang.org/x/net/context"
)

type ActivityServiceImpl struct{}

func (h ActivityServiceImpl) GetById(ctx context.Context, activityPb *Activity) (*Activity, error) {
	activityRes := new(Activity)
	ActivityModel := new(models.ActivityModel)
	activityData := ActivityModel.GetById(activityPb.Id)
	copier.Copy(activityRes, activityData)
	return activityRes, nil
}

func (h ActivityServiceImpl) List(ctx context.Context, in *Query) (*Activitylist, error) {
	return nil, nil
}

func (h ActivityServiceImpl) Insert(ctx context.Context, activityPb *Activity) (*Activity, error) {
	return nil, nil
}
