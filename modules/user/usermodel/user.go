package usermodel

import (
	"errors"
	"go-simple-api/common"
	"go-simple-api/component/tokenprovider"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" db:"email"`
	Password        string `json:"-" db:"password"`
	Salt            string `json:"-" db:"salt"`
	LastName        string `json:"last_name" db:"last_name"`
	FirstName       string `json:"first_name" db:"first_name"`
	Phone           string `json:"phone" db:"phone"`
	Role            string `json:"role" db:"roles"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	Phone           string `json:"phone"`
	Role            string `json:"-"`
	Salt            string `json:"-"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
