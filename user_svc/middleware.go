package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const svcName = "user"

type LogMiddleware struct {
	next UserService
}

func NewLogMiddleware(next UserService) UserService {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CreateUser(ctx context.Context, params *UserDBCreateParams) (u *types.User, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"user_email": fmt.Sprintf("%v", u.Email),
			"user_id": fmt.Sprintf("%v", u.ID),
		}).Info("create user")
	}(time.Now())
	return l.next.CreateUser(ctx, params)
}

func (l *LogMiddleware) ReadUser(ctx context.Context, filter *types.ReadQuery) (u *types.User, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"user_email": fmt.Sprintf("%v", u.Email),
			"user_id": fmt.Sprintf("%v", u.ID),
		}).Info("read user")
	}(time.Now())
	return l.next.ReadUser(ctx, filter)
}

func (l *LogMiddleware) UpdateUser(ctx context.Context, filter *types.ReadQuery, update *UserDBUpdateParams) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"update": fmt.Sprintf("%v", filter),
		}).Info("update user")
	}(time.Now())
	return l.next.UpdateUser(ctx, filter, update)
}

func (l *LogMiddleware) DeleteUser(ctx context.Context, id string) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"deleted_user_id": id,
		}).Info("delete user")
	}(time.Now())
	return l.next.DeleteUser(ctx, id)
}

func (l *LogMiddleware)  ListUser(ctx context.Context) (users []*types.User, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"user_read_count": len(users),
		}).Info("delete user")
	}(time.Now())
	return l.next.ListUser(ctx)
}