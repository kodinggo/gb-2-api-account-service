package usecase

import "account-service/src/model"

type accountRepository struct {
	accountRepository model.AccountRepository
}

func NewAccountUsecase(accountRepository model.AccountRepository) model.AccountUsecase {

}
