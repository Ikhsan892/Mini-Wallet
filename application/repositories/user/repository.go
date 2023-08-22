package user

import (
	"context"
	"github.com/stretchr/testify/mock"
	"wallet/application/entity"
)

type UserRepository interface {
	CreateUserAccount(ctx context.Context, userId string) (entity.User, error)
	GetUserAccountFromToken(ctx context.Context, token string) (entity.User, error)
	IsTokenValid(ctx context.Context, token string) bool
}

type UserMockrepository struct {
	mock.Mock
}

func (u *UserMockrepository) CreateUserAccount(ctx context.Context, userId string) (entity.User, error) {
	param := u.Called(ctx, userId)

	return param.Get(0).(entity.User), param.Error(1)
}

func (u *UserMockrepository) IsTokenValid(ctx context.Context, token string) bool {
	param := u.Called(ctx, token)

	return param.Get(0).(bool)
}

func (u *UserMockrepository) GetUserAccountFromToken(ctx context.Context, token string) (entity.User, error) {
	param := u.Called(ctx, token)

	return param.Get(0).(entity.User), param.Error(1)
}
