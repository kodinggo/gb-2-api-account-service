package usecase

import (
	"context"
	"errors"

	"github.com/kodinggo/gb-2-api-account-service/src/helper"
	"github.com/kodinggo/gb-2-api-account-service/src/model"

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
func (u *accountUsecase) FindByID(ctx context.Context, data model.Account, id int64) (*model.Account, error) {
	account, err := u.accountRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("not found")
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

func (u *accountUsecase) FindByIDs(ctx context.Context, ids []int64) ([]*model.Account, error) {
	accounts, err := u.accountRepository.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	if accounts == nil {
		return nil, errors.New("not found")
	}
	return accounts, nil
}
