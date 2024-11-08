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

func (r *accountRepository) RegisterNewUser(ctx context.Context, data model.Account) (newAccount *model.Account, err error) {
	now := time.Now().UTC()

	result, err := sq.Insert("accounts").
		Columns("fullname", "sort_bio", "gender", "picture_url", "username", "email", "password", "role", "created_at", "updated_at").
		Values(data.Fullname, data.SortBio, data.Gender, data.PictureUrl, data.Username, data.Email, data.Password, data.Role, now, now).
		RunWith(r.db).
		ExecContext(ctx)

	if err != nil {
		logrus.WithField("data", data).Error(err)
	}

	lastInsertId, _ := result.LastInsertId()

	newAccount = &data
	newAccount.ID = lastInsertId
	newAccount.CreatedAt = now

	return
}
