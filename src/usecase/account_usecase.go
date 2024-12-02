package usecase

import (
	"account-service/src/helper"
	"account-service/src/model"
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type accountUsecase struct {
	accountRepository model.AccountRepository
}

func NewAccountUsecase(accountRepository model.AccountRepository) model.AccountUsecase {
	return &accountUsecase{accountRepository: accountRepository}
}

func (account *accountUsecase) Create(ctx context.Context, data model.Register) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data,
	})

	passwordHashed, err := helper.HashRequestPassword(data.Password)
	if err != nil {
		logger.Error(err)
		return
	}

	newAccount, err := account.accountRepository.Store(ctx, model.Account{
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

	user := u.accountRepository.FindByEmail(ctx, data.Email)
	if user == nil {
		err = errors.New("wrong email or password")
		return
	}

	if !helper.CheckPasswordHash(data.Password, user.Password) {
		err = errors.New("missmatch password")
		return
	}

	token, err = helper.GenerateToken(user.ID)

	if err != nil {
		logger.Error(err)
	}

	return
}

func (u *accountUsecase) FindById(ctx context.Context, data model.Account, id int64) (*model.Account, error) {
	account, err := u.accountRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, fmt.Errorf("account with id %d not found", id)
	}

	return account, nil
}

func (u *accountUsecase) Update(ctx context.Context, data model.Account, id int64) (*model.Account, error) {
	logger := logrus.WithFields(logrus.Fields{
		"email": data.Email,
		"id":    id,
	})

	updatedAccount, err := u.accountRepository.Update(ctx, data, id)
	if err != nil {
		logger.Error("Failed to update account: ", err)
		return nil, err
	}

	logger.Info("Account updated successfully")
	return updatedAccount, nil
}
