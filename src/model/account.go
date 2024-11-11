package model

import (
	"context"
	"mime/multipart"
	"time"
)

type AccountRepository interface {
	RegisterNewAccountToDatabase(ctx context.Context, data Account) (*Account, error)
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
	CreateNewAccountData(ctx context.Context, data Register, fileHeader *multipart.FileHeader) (token string, err error)
	Login(ctx context.Context, data Login) (token string, err error)
}

type Register struct {
	Fullname   string `json:"fullname" form:"fullname"`
	SortBio    string `json:"sort_bio" form:"sort_bio"`
	Gender     Gender `json:"gender" form:"gender"`
	PictureUrl string `json:"picture" form:"picture"`
	Username   string `json:"username" form:"username"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       Role   `json:"role" form:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
