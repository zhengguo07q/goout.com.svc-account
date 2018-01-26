package account

import (
	"github.com/jinzhu/gorm"
	"service-account/account"
)

var DB *gorm.DB

type Account struct {
	ID       int `gorm:"primary_key:id"`
	Username string
	Email    string
	Phone    string
	Password string
	Birthday time.Time
	state    AccountState
}

func (account *Account) Insert() error {
	return DB.Create(account).Error
}

func (account *Account) Update() error {
	return DB.Save(account).Error
}

func (account *Account) FirstOrCreate() (*Account, error) {
	err := DB.FirstOrCreate(account, "github_login_id = ?", account.GithubLoginId).Error
	return user, err
}

func (account *Account) UpdateEmail(email string) error {
	if len(email) > 0 {
		return DB.Model(account).Update("email", email).Error
	} else {
		return DB.Model(account).Update("email", gorm.Expr("NULL")).Error
	}
}

func (account *Account) Lock() error {
	return DB.Model(account).Update(map[string]interface{}{
		"lock_state": account.LockState,
	}).Error
}

func GetAccount(id interface{}) (*Account, error) {
	var account Account
	err := DB.First(&account, id).Error
	return &account, err
}

func GetAccountByKey(key string) (*Account, error) {
	var account Account
	err := DB.First(&user, "email = ?", key).Error
	return &account, err
}

func ListAccounts() ([]*Account, error) {
	var accounts []*Account
	err := DB.Find(&accounts, "is_admin = ?", false).Error
	return accounts, err
}
