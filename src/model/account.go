package model

import (
	"context"
	"time"
)

type AccountRepository interface {
	Store(ctx context.Context, data Account) (*Account, error)
	FindByEmail(ctx context.Context, email string) *Login
	FindByID(ctx context.Context, id int64) (*Account, error)
	//TODO : UpdateAccount & FindAccount By Id

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
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	Role       Role      `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type AccountUsecase interface {
	Create(ctx context.Context, data Register) (token string, err error)
	Login(ctx context.Context, data Login) (token string, err error)
	FindByID(ctx context.Context, id int64) (*Account, error)
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
