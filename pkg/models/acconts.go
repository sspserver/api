package models

import (
	gogosql "github.com/geniusrabbit/gosql/gorm"

	"github.com/geniusrabbit/blaze-api/repository/account"
	accountModels "github.com/geniusrabbit/blaze-api/repository/account/models"
	"github.com/geniusrabbit/blaze-api/repository/user"
	userModels "github.com/geniusrabbit/blaze-api/repository/user/models"
)

type Contact struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	IsPrimary bool   `json:"is_primary,omitempty"`
}

// Account is the bundled default for example/api: Base + consumer profile trait.
type Account struct {
	accountModels.AccountBase

	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`

	CountryCode      string `gorm:"column:country_code;type:varchar(2)" json:"country_code"`
	City             string `gorm:"column:city;type:varchar(255)" json:"city"`
	ZipCode          string `gorm:"column:zip_code;type:varchar(20)" json:"zip_code"`
	Address          string `gorm:"column:address;type:varchar(255)" json:"address"`
	Phone            string `gorm:"column:phone;type:varchar(20)" json:"phone"`
	VATNumber        string `gorm:"column:vat_number;type:varchar(50)" json:"vat_number"`
	CompanyRegNumber string `gorm:"column:company_reg_number;type:varchar(50)" json:"company_reg_number"`

	Contacts gogosql.NullableJSONArray[Contact] `gorm:"column:contacts;type:json" json:"contacts"`
}

// NewWithIDs returns a new account instance with ID and admin user IDs set.
func (a *Account) NewWithIDs(id uint64, adminUserIDs ...uint64) account.Model {
	return &Account{AccountBase: accountModels.AccountBase{ID: id, Admins: adminUserIDs}}
}

// User is the example consumer user model (Base + Email + Password traits).
type User struct {
	userModels.UserBase
	userModels.UserEmail
	userModels.UserPassword
}

func (u *User) NewWithID(id uint64) user.Model {
	return &User{UserBase: userModels.UserBase{ID: id}}
}

// Anonymous is the example anonymous session user placeholder.
var Anonymous = User{UserBase: userModels.UserBase{ID: 0}}

// AccountMember is the bundled member type for example/api.
type AccountMember = account.Member[*User, *Account]
