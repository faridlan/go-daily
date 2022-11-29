package service

import (
	"context"

	"github.com/faridlan/daily/test/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreate) web.UserResponse
	Update(ctx context.Context, request web.UserUpdate) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
