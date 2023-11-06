package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"errors"
)

type AccountService interface {
	CreateAccount(ctx context.Context, params *AccountDBCreateParams) (*types.Account, error)
	ReadAccount(ctx context.Context, filter *types.ReadQuery) (*types.Account, error)
	UpdateAccount(ctx context.Context, filter *types.ReadQuery, update *AccountDBUpdateParams) error
	DeleteAccount(ctx context.Context, id string) error
	ListAccount(ctx context.Context) ([]*types.Account, error)
}

type AccountServiceImpl struct {
	store Store
}

func NewAccountService(s Store) AccountService {
	return &AccountServiceImpl{
		store: s,
	}
}

func (s *AccountServiceImpl) CreateAccount(ctx context.Context, params *AccountDBCreateParams) (*types.Account, error) {
	if errList := params.Validate(); len(errList) > 0 {
		return nil, errors.New("error validating user params")
	}
	acc := &types.Account{
		UserID: params.UserID,
		StripeCustomerID: params.StripeCustomerID,
		StripeSubscriptionID: params.StripeSubscriptionID,
		SubscriptionStatus: params.SubscriptionStatus,
		Plan: params.Plan,
	}
	a, err := s.store.Insert(ctx, acc)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AccountServiceImpl) ReadAccount(ctx context.Context, filter *types.ReadQuery) (*types.Account, error) {
	dbFilter := types.DBFilter{filter.By: filter.Item}
	a, err := s.store.Get(ctx, dbFilter)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AccountServiceImpl) UpdateAccount(ctx context.Context, filter *types.ReadQuery, update *AccountDBUpdateParams) error {
	dbFilter := types.DBFilter{filter.By: filter.Item}
	if err := s.store.Update(ctx, dbFilter, update); err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) DeleteAccount(ctx context.Context, id string) error {
	if err := s.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *AccountServiceImpl) ListAccount(ctx context.Context) ([]*types.Account, error) {
	a, err := s.store.List(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}
