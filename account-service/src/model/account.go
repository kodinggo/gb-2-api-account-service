package model

import (
	"context"
	"time"
)

type AccountRepository interface {
	// FindByEmail(ctx context.Context, email string) (*Account, error)
	RegisterNewUser(ctx context.Context, data Account) (*Account, error)
}

type AccountUsecase interface {
	Register(ctx context.Context, data Register) (string error)
	Login(ctx context.Context, data Login) (string error)
}

type Register struct {
	Fullname   string `json:"fullname"`
	SortBio    string `json:"string_bio"`
	Gender     Gender `json:"gender"`
	PictureUrl string `json:"picture_url"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Gender string
type Role string

const (
	MALE    Gender = "Male"
	FEMALE  Gender = "Female"
	OTHERES Gender = "Others"
)

const (
	ADMIN  Role = "Admin"
	MEMBER Role = "Member"
)

type Account struct {
	ID         int64     `json:"id"`
	Fullname   string    `json:"fullname"`
	SortBio    string    `json:"string_bio"`
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
