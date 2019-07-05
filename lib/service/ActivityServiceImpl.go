package service

import (
	"chs/models"
	"context"
	"github.com/jinzhu/copier"
)

type ActivityServiceImpl struct{}

func (h ActivityServiceImpl) GetById(ctx context.Context, activity *Activity) (*Activity, error) {
	activityRes := new(Activity)
	ActivityModel := new(models.ActivityModel)
	activityData := ActivityModel.GetById(activity.Id)
	copier.Copy(activityRes, activityData)
	return activityRes, nil
}

func (h ActivityServiceImpl) List(ctx context.Context, in *ActivityQuery) (*Activitylist, error) {
	return nil, nil
}

func (h ActivityServiceImpl) Insert(ctx context.Context, activityPb *Activity) (*Activity, error) {
	return nil, nil
}
