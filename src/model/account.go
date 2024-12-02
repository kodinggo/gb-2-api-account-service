package model

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ContextAuthKey string

const BearerAuthKey ContextAuthKey = "BearerAuth"

type AccountRepository interface {
	Store(ctx context.Context, data Account) (*Account, error)
	FindByEmail(ctx context.Context, email string) *Login
	FindById(ctx context.Context, id int64) (*Account, error)
	Update(ctx context.Context, account Account, id int64) (*Account, error)
}

type Gender string
type Role string

const (
	MALE   Gender = "male"
	FEMALE Gender = "female"
	OTHERS Gender = "others"
)

const (
	ADMIN  Role = "admin"
	MEMBER Role = "member"
)

type Account struct {
	ID         int64     `json:"id"`
	Fullname   string    `json:"fullname"`
	SortBio    string    `json:"sort_bio"`
	Gender     Gender    `json:"gender"`
	PictureUrl string    `json:"picture_url"`
	Username   string    `json:"-"`
	Email      string    `json:"-"`
	Password   string    `json:"-"`
	Role       Role      `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type AccountUsecase interface {
	Create(ctx context.Context, data Register) (token string, err error)
	Login(ctx context.Context, data Login) (token string, err error)
	FindById(ctx context.Context, data Account, id int64) (*Account, error)
	Update(ctx context.Context, data Account, id int64) (*Account, error)
}

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims

}

type Register struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
