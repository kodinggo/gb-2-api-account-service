package usecase

import (
	"account-service/src/helper"
	"account-service/src/model"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/sirupsen/logrus"
)

type accountUsecase struct {
	accountRepository model.AccountRepository
}

func NewAccountUsecase(accountRepository model.AccountRepository) model.AccountUsecase {
	return &accountUsecase{accountRepository: accountRepository}
}

func (account *accountUsecase) CreateNewAccountData(ctx context.Context, data model.Register, fileHeader *multipart.FileHeader) (token string, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"data": data,
	})

	logrus.Infof("test ini dikasih dari usecase: %s", data)

	imageUploaded, err := helper.SaveUploadedFile(fileHeader, "upload/picture/account")
	if err != nil {
		logger.Error(err)
		return "", fmt.Errorf("failed to save uploaded file: %w", err)
	}
	data.PictureUrl = imageUploaded
	passwordHashed, err := helper.HashRequestPassword(data.Password)

	if err != nil {
		logger.Error(err)
		return
	}

	newAccount, err := account.accountRepository.RegisterNewAccountToDatabase(ctx, model.Account{
		Fullname:   data.Fullname,
		SortBio:    data.SortBio,
		Gender:     data.Gender,
		PictureUrl: data.PictureUrl,
		Username:   data.Username,
		Email:      data.Email,
		Password:   passwordHashed,
		Role:       data.Role,
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
