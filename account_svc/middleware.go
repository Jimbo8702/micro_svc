package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const svcName = "account"

type LogMiddleware struct {
	next AccountService
}

func NewLogMiddleware(next AccountService) AccountService {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CreateAccount(ctx context.Context, params *AccountDBCreateParams) (a *types.Account, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"account_user_id": fmt.Sprintf("%v", a.UserID),
			"account_id": fmt.Sprintf("%v", a.ID),
		}).Info("create account")
	}(time.Now())
	return l.next.CreateAccount(ctx, params)
}

func (l *LogMiddleware) ReadAccount(ctx context.Context, filter *types.ReadQuery) (a *types.Account, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"account_id": fmt.Sprintf("%v", a.ID),
		}).Info("read account")
	}(time.Now())
	return l.next.ReadAccount(ctx, filter)
}

func (l *LogMiddleware) UpdateAccount(ctx context.Context, filter *types.ReadQuery, update *AccountDBUpdateParams) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"update": fmt.Sprintf("%v", filter),
		}).Info("update account")
	}(time.Now())
	return l.next.UpdateAccount(ctx, filter, update)
}

func (l *LogMiddleware) DeleteAccount(ctx context.Context, id string) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"deleted_account_id": id,
		}).Info("delete account")
	}(time.Now())
	return l.next.DeleteAccount(ctx, id)
}

func (l *LogMiddleware)  ListAccount(ctx context.Context) (accounts []*types.Account, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"service_name": svcName,
			"took": time.Since(start),
			"err": err,
			"account_read_count": len(accounts),
		}).Info("delete account")
	}(time.Now())
	return l.next.ListAccount(ctx)
}