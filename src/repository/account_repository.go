package repository

import (
	"account-service/src/model"
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) model.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) RegisterNewAccountToDatabase(ctx context.Context, data model.Account) (newAccount *model.Account, err error) {
	now := time.Now().UTC()

	result, err := sq.Insert("accounts").
		Columns("fullname", "sort_bio", "gender", "picture_url", "username", "email", "password", "role", "created_at", "updated_at").
		Values(data.Fullname, data.SortBio, data.Gender, data.PictureUrl, data.Username, data.Email, data.Password, data.Role, now, now).
		RunWith(r.db).
		ExecContext(ctx)

	if err != nil {
		logrus.WithField("data", data).Error(err)
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		logrus.Error("Error getting last insert ID:", err)
	} else {
		logrus.Infof("Last Insert ID: %d", lastInsertId)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logrus.Error("Error getting rows affected:", err)
	} else {
		logrus.Infof("Rows Affected: %d", rowsAffected)
	}

	newAccount = &data
	newAccount.ID = lastInsertId
	newAccount.CreatedAt = now

	return
}

// TODO : Login and FindById
