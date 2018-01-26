package account

import (
	"crypto/md5"
	"errors"

	"strings"
	"svc_account"
	"time"
)

const PasswordLen = 8



type AccountState int

//在线状态
const (
	ASOffline AccountState = iota
	ASHide
	ASOnline
)

func (as AccountState) String() string {
	switch as {
	case ASOffline:
		return "Offline"
	case ASHide:
		return "ASHide"
	case ASOnline:
		return "ASOnline"
	}
}

type Service interface {
	//创建账户
	CreateAccount(username string, password string, birthday time.Time, email string, phone string) (int, error)
	//通过登录KEY，密码进行登录(登录key可以为以下几种类型， username, phone, email)
	Login(loginKey string, password string) (Account, error)
	//登出
	Logout(id int) error
	//改变密码
	ChangePassword(id int, oldPassword string, newPassword string) error
	//告知状态
	SetOnline(id int) error
}

type service struct{}

//创建账户需要改写
func (service) CreateAccount(username string, email string, phone string, password string, birthday time.Time) (int, error) {
	if username == "" || email == "" || phone == "" {
		return 0, ErrAccountNull
	}
	if password == "" {
		return 0, ErrPasswordNull
	}
	if len(password) < PasswordLen {
		return 0, ErrPasswordNotEnoughLen
	}
	passwordMd5 := md5.Sum(password)

	account := &Account{Username: username}
	_, err := DB.Create(account)
	if err != nil {
		return 0, ErrUnknow
	}
	return 1, nil
}

func (service) Login(loginKey string, password string) (Account, error) {
	if loginKey == "" {
		return 0, ErrAccountNull
	}
	if password == "" {
		return 0, ErrPasswordNull
	}

	var account Account
	_, err := DB.First(&account, "username = ? or email = ? or phone = ?")
	if err != nil {
		return 0, ErrUnknow
	}

	return account, nil
}

func (service) Logout(id int) error {
	var account Account
	_, err := DB.First(&account, id)
	if err != nil {
		return 0, ErrUnknow
	}

}

func (service) ChangePassword(id int, oldPassword string, newPassword string) error {

}


