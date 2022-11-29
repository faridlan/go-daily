package service

import (
	"context"
	"database/sql"

	"github.com/faridlan/daily/test/helper"
	"github.com/faridlan/daily/test/model/domain"
	"github.com/faridlan/daily/test/model/web"
	"github.com/faridlan/daily/test/repository"
)

type UserServiceImpl struct {
	DB       sql.DB
	UserRepo repository.UserRepository
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreate) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Password,
		Password: request.Password,
	}

	user = service.UserRepo.Create(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdate) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(err)
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepo.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	service.UserRepo.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepo.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepo.FindAll(ctx, tx)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponses(users)
}
