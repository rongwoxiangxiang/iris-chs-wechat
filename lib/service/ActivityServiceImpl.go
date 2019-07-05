package service

import (
	"context"
)

type ActivityServiceImpl struct{}

func (ActivityServiceImpl) GetById(context.Context, *Activity) (*Activity, error) {
	panic("implement me")
}

func (ActivityServiceImpl) List(context.Context, *Query) (*Activitylist, error) {
	panic("implement me")
}

func (ActivityServiceImpl) Insert(context.Context, *Activity) (*Activity, error) {
	panic("implement me")
}

//func (h ActivityServiceImpl) GetById(ctx context.Context, activity *Activity) (*Activity, error) {
//	activityRes := new(Activity)
//	ActivityModel := new(models.ActivityModel)
//	activityData := ActivityModel.GetById(activity.Id)
//	copier.Copy(activityRes, activityData)
//	return activityRes, nil
//}
//
//func (h ActivityServiceImpl) List(ctx context.Context, in *Query) (*Activitylist, error) {
//	return nil, nil
//}
//
//func (h ActivityServiceImpl) Insert(ctx context.Context, activityPb *Activity) (*Activity, error) {
//	return nil, nil
//}
