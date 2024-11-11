package usecase

import (
	"account-service/src/helper"
	"account-service/src/model"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

type accountUsecase struct {
	accountRepository model.AccountRepository
}

func NewAccountUsecase(accountRepository model.AccountRepository) model.AccountUsecase {
	return &accountUsecase{accountRepository: accountRepository}
}

func (account *accountUsecase) CreateAccount(ctx context.Context, data model.Register) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data,
	})

	passwordHashed, err := helper.HashRequestPassword(data.Password)

	if err != nil {
		logger.Error(err)
		return
	}

	newAccount, err := account.accountRepository.StoreAccount(ctx, model.Account{
		Username: data.Username,
		Email:    data.Email,
		Password: passwordHashed,
	})

	if err != nil {
		logger.Error(err)

		return
	}

	acceesToken, err := helper.GenerateToken(newAccount.ID)

	if err != nil {
		logger.Error(err)
		return
	}

	return acceesToken, nil
}

func (u *accountUsecase) Login(ctx context.Context, data model.Login) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data.Email,
	})

	user, err := u.accountRepository.FindByEmail(ctx, data.Email)

	if err != nil {
		logger.Error(err)
		return
	}

	if !helper.CheckPasswordHash(data.Password, user.Password) {
		err = errors.New("incorrect password")
		return
	}

	token, err = helper.GenerateToken(user.ID)

	if err != nil {
		logger.Error(err)
	}

	return
}
