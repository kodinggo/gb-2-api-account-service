package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kodinggo/gb-2-api-account-service/src/model"

	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) model.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Store(ctx context.Context, data model.Account) (newAccount *model.Account, err error) {
	now := time.Now().UTC()

	result, err := sq.Insert("accounts").
		Columns("username", "email", "password", "created_at", "updated_at").
		Values(data.Username, data.Email, data.Password, now, now).
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

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (account *model.Login) {
	row := sq.Select("id", "email", "password").
		From("accounts").
		Where(sq.Eq{"email": email}).
		RunWith(r.db).
		QueryRowContext(ctx)

	var data model.Login
	err := row.Scan(
		&data.ID,
		&data.Email,
		&data.Password,
	)
	if err != nil {
		return
	}

	return &data
}

func (r *accountRepository) FindByID(ctx context.Context, id int64) (*model.Account, error) {
	row := sq.Select("id", "fullname", "sort_bio", "gender", "picture_url", "username", "role").
		From("accounts").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		QueryRowContext(ctx)

	var fullname, sortBio, gender, pictureUrl sql.NullString

	var data model.Account
	err := row.Scan(
		&data.ID,
		&fullname,
		&sortBio,
		&gender,
		&pictureUrl,
		&data.Username,
		&data.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account with id %d: %w", id, err)
	}
	data.Fullname = fullname.String
	data.SortBio = sortBio.String
	data.Gender = model.Gender(gender.String)
	data.PictureUrl = pictureUrl.String
	data.Role = model.Role(data.Role)

	return &data, nil
}

func (r *accountRepository) Update(ctx context.Context, account model.Account, id int64) (*model.Account, error) {
	_, err := sq.Update("accounts").
		Set("fullname", account.Fullname).
		Set("sort_bio", account.SortBio).
		Set("gender", account.Gender).
		Set("picture_url", account.PictureUrl).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		ExecContext(ctx)

	if err != nil {
		return nil, err
	}

	updatedAccount, err := r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return updatedAccount, nil

}

func (r *accountRepository) FindByIDs(ctx context.Context, ids []int64) ([]*model.Account, error) {
	rows, err := sq.Select("id", "fullname", "sort_bio", "gender", "picture_url", "username", "email").
		From("accounts").
		Where(sq.Eq{"id": ids}).
		RunWith(r.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*model.Account
	for rows.Next() {
		var (
			fullname   sql.NullString
			sortBio    sql.NullString
			gender     sql.NullString
			pictureURL sql.NullString
			username   sql.NullString
			email      sql.NullString
		)

		account := &model.Account{}
		err := rows.Scan(
			&account.ID,
			&fullname,
			&sortBio,
			&gender,
			&pictureURL,
			&username,
			&email,
		)
		if err != nil {
			return nil, err
		}

		account.Fullname = fullname.String
		account.SortBio = sortBio.String
		account.Gender = model.Gender(gender.String)
		account.PictureUrl = pictureURL.String
		account.Username = username.String
		account.Email = email.String

		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
