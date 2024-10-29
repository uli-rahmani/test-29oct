package usecase

import (
	"context"
	"errors"
	"fmt"
	"test/domain"
	"test/domain/general"
	"test/repo/db"
	"test/utils"
)

type UserUsecase struct {
	conf     *general.SectionService
	userRepo db.UserRepoItf
}

func newUserUsecase(conf *general.SectionService, dbRepo db.DBRepo) UserUsecase {
	return UserUsecase{
		userRepo: dbRepo.User,
		conf:     conf,
	}
}

type UserUsecaseIntf interface {
	// RegisterUser(ctx context.Context, data domain.RegisterUserRequest) (domain.RegisterUserResponse, error)
	Login(ctx context.Context, data domain.LoginRequest) (domain.LoginResponse, string, error)
}

func (uu UserUsecase) Login(ctx context.Context, data domain.LoginRequest) (domain.LoginResponse, string, error) {
	var resp domain.LoginResponse

	userData, err := uu.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return resp, "failed to check user data", err
	}

	if userData == nil {
		return resp, "user not registered", errors.New("users not found")
	}

	passwordValid, err := utils.ComparePassword(userData.Password, data.Password)
	if err != nil {
		return resp, "failed to check user data", err
	}

	if !passwordValid {
		return resp, "password not match", errors.New("password wrong")
	}

	session, err := utils.GetEncrypt([]byte(uu.conf.Authorization.JWT.AccessTokenSecretKey), fmt.Sprintf("%d", userData.ID))
	if err != nil {
		return resp, "failed to check user data", err
	}

	jwtData, err := utils.GenerateJWT(session, nil)
	if err != nil {
		return resp, "failed to check user data", err
	}

	resp = domain.LoginResponse{
		Token: jwtData,
	}

	return resp, "login success", nil
}
