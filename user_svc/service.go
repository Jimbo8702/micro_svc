package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"errors"
)

type UserService interface {
	CreateUser(ctx context.Context, params *UserDBCreateParams) (*types.User, error)
	ReadUser(ctx context.Context, filter *types.ReadQuery) (*types.User, error)
	UpdateUser(ctx context.Context, filter *types.ReadQuery, update *UserDBUpdateParams) error
	DeleteUser(ctx context.Context, id string) error
	ListUser(ctx context.Context) ([]*types.User, error)
}

type UserServiceImpl struct {
	store Store
}

func NewUserService(s Store) UserService {
	return &UserServiceImpl{
		store: s,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, params *UserDBCreateParams) (*types.User, error) {
	if errList := params.Validate(); len(errList) > 0 {
		return nil, errors.New("error validating user params")
	}
	ps, err := NewPassword(params.Password)
	if err != nil {
		return nil, errors.New("error generating new password")
	}
	user := &types.User{
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		EncryptedPassword: ps,
	}
	u, err := s.store.Insert(ctx, user)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserServiceImpl) ReadUser(ctx context.Context, filter *types.ReadQuery) (*types.User, error) {
	dbFilter := types.DBFilter{filter.By: filter.Item}
	u, err := s.store.Get(ctx, dbFilter)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserServiceImpl) ListUser(ctx context.Context) ([]*types.User, error) {
	u, err := s.store.List(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, filter *types.ReadQuery, update *UserDBUpdateParams) error {
	dbFilter := types.DBFilter{filter.By: filter.Item}
	if update.Password != "" {
		ps, err := NewPassword(update.Password);
		if err != nil {
			return err
		}
		update.Password = ps
	}
	if err := s.store.Update(ctx, dbFilter, update); err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, id string) error {
	if err := s.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}