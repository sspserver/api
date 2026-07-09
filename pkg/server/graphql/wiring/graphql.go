package wiring

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	gogosql "github.com/geniusrabbit/gosql/gorm"

	pkgModels "github.com/geniusrabbit/blaze-api/pkg/models"
	"github.com/geniusrabbit/blaze-api/repository/user"
	basemodels "github.com/geniusrabbit/blaze-api/server/graphql/models"
	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// AccountToGraphQL maps example domain Account to extended GraphQL Account.
func AccountToGraphQL(acc *models.Account) *gqlmodels.Account {
	if acc == nil {
		return nil
	}
	return &gqlmodels.Account{
		ID:               acc.GetID(),
		Status:           basemodels.ApproveStatusFrom(acc.GetApprove()),
		Name:             acc.Name,
		Description:      acc.Description,
		CountryCode:      gocast.IfThen(acc.CountryCode != "", &acc.CountryCode, nil),
		City:             gocast.IfThen(acc.City != "", &acc.City, nil),
		ZipCode:          gocast.IfThen(acc.ZipCode != "", &acc.ZipCode, nil),
		Address:          gocast.IfThen(acc.Address != "", &acc.Address, nil),
		Phone:            gocast.IfThen(acc.Phone != "", &acc.Phone, nil),
		VatNumber:        gocast.IfThen(acc.VATNumber != "", &acc.VATNumber, nil),
		CompanyRegNumber: gocast.IfThen(acc.CompanyRegNumber != "", &acc.CompanyRegNumber, nil),
		CreatedAt:        acc.GetCreatedAt(),
		UpdatedAt:        acc.GetUpdatedAt(),
		Contacts: xtypes.SliceApply(acc.Contacts, func(c models.Contact) *gqlmodels.Contact {
			return &gqlmodels.Contact{
				Type:      c.Type,
				Value:     c.Value,
				IsPrimary: gocast.IfThen(c.IsPrimary, &c.IsPrimary, nil),
			}
		}),
	}
}

// FillAccountFromCreateInput copies account create input into domain Account.
func FillAccountFromCreateInput(dest *models.Account, input *gqlmodels.AccountCreateInput, appStatus ...pkgModels.ApproveStatus) *models.Account {
	if dest == nil || input == nil {
		return dest
	}

	status := pkgModels.UndefinedApproveStatus
	if len(appStatus) > 0 {
		status = appStatus[0]
	} else if input.Status != nil {
		status = input.Status.ModelStatus()
	}
	dest.SetApprove(status)

	dest.Name = input.Name
	dest.Description = gocast.PtrAsValue(input.Description, dest.Description)

	dest.CountryCode = gocast.PtrAsValue(input.CountryCode, dest.CountryCode)
	dest.City = gocast.PtrAsValue(input.City, dest.City)
	dest.ZipCode = gocast.PtrAsValue(input.ZipCode, dest.ZipCode)
	dest.Address = gocast.PtrAsValue(input.Address, dest.Address)
	dest.Phone = gocast.PtrAsValue(input.Phone, dest.Phone)
	dest.VATNumber = gocast.PtrAsValue(input.VatNumber, dest.VATNumber)
	dest.CompanyRegNumber = gocast.PtrAsValue(input.CompanyRegNumber, dest.CompanyRegNumber)

	dest.Contacts = gogosql.NullableJSONArray[models.Contact](
		xtypes.SliceApply(input.Contacts, func(c *gqlmodels.ContactInput) models.Contact {
			return models.Contact{
				Type:      c.Type,
				Value:     c.Value,
				IsPrimary: gocast.PtrAsValue(c.IsPrimary, false),
			}
		}),
	)

	return dest
}

// FillAccountFromUpdateInput applies account update input to a domain Account.
func FillAccountFromUpdateInput(dest *models.Account, input *gqlmodels.AccountUpdateInput, appStatus ...pkgModels.ApproveStatus) *models.Account {
	if dest == nil {
		return dest
	}
	if len(appStatus) > 0 {
		dest.SetApprove(appStatus[0])
	} else if input != nil && input.Status != nil {
		dest.SetApprove(input.Status.ModelStatus())
	}
	if input != nil {
		dest.Name = gocast.PtrAsValue(input.Name, dest.Name)
		dest.Description = gocast.PtrAsValue(input.Description, dest.Description)

		dest.CountryCode = gocast.PtrAsValue(input.CountryCode, dest.CountryCode)
		dest.City = gocast.PtrAsValue(input.City, dest.City)
		dest.ZipCode = gocast.PtrAsValue(input.ZipCode, dest.ZipCode)
		dest.Address = gocast.PtrAsValue(input.Address, dest.Address)
		dest.Phone = gocast.PtrAsValue(input.Phone, dest.Phone)
		dest.VATNumber = gocast.PtrAsValue(input.VatNumber, dest.VATNumber)
		dest.CompanyRegNumber = gocast.PtrAsValue(input.CompanyRegNumber, dest.CompanyRegNumber)
		dest.Contacts = gogosql.NullableJSONArray[models.Contact](
			xtypes.SliceApply(input.Contacts, func(c *gqlmodels.ContactInput) models.Contact {
				return models.Contact{
					Type:      c.Type,
					Value:     c.Value,
					IsPrimary: gocast.PtrAsValue(c.IsPrimary, false),
				}
			}),
		)
	}
	return dest
}

// UserToGraphQL maps example domain User to base GraphQL User.
func UserToGraphQL(u *models.User) *gqlmodels.User {
	if u == nil {
		return nil
	}
	return &gqlmodels.User{
		ID:        u.GetID(),
		Email:     u.GetEmail(),
		Status:    basemodels.ApproveStatusFrom(u.GetApprove()),
		CreatedAt: u.GetCreatedAt(),
		UpdatedAt: u.GetUpdatedAt(),
	}
}

// UserToGraphQLPtr maps example domain User to base GraphQL User pointer (account/member payloads).
func UserToGraphQLPtr(u *models.User) *gqlmodels.User {
	if u == nil {
		return nil
	}
	return UserToGraphQL(u)
}

// UserFromCreateInput builds a new domain User from GraphQL create input.
func UserFromCreateInput(input *gqlmodels.UserCreateInput, appStatus ...pkgModels.ApproveStatus) *models.User {
	if input == nil {
		return nil
	}
	u := new(models.User)
	var status pkgModels.ApproveStatus
	if len(appStatus) > 0 {
		status = appStatus[0]
	}
	u.SetApprove(status)
	u.SetEmail(input.Email)
	return u
}

// FillUserFromInput merges GraphQL update input into an existing domain User.
func UserFromUpdateInput(input *gqlmodels.UserUpdateInput, dest *models.User, appStatus ...pkgModels.ApproveStatus) *models.User {
	if dest == nil {
		return nil
	}
	if len(appStatus) > 0 {
		dest.SetApprove(appStatus[0])
	} else if input != nil && input.Status != nil {
		dest.SetApprove(input.Status.ModelStatus())
	}
	if input != nil && input.Email != nil {
		dest.SetEmail(*input.Email)
	}
	return dest
}

// UserListFilterMapper converts extended user list filter to domain query option.
func UserListFilterMapper(fl *gqlmodels.UserListFilter) user.QOption {
	if fl == nil {
		return nil
	}
	return &user.ListFilter{
		FilterBase: user.FilterBase{
			ID: fl.ID,
		},
		FilterEmail: user.FilterEmail{
			Emails: fl.Emails,
		},
	}
}

// UserListOrderMapper converts extended user list order to domain query option.
func UserListOrderMapper(ord *gqlmodels.UserListOrder) user.QOption {
	if ord == nil {
		return nil
	}
	return &user.ListOrder{
		OrderBase: user.OrderBase{
			ID:        ord.ID.AsOrder(),
			Status:    ord.Status.AsOrder(),
			CreatedAt: ord.CreatedAt.AsOrder(),
			UpdatedAt: ord.UpdatedAt.AsOrder(),
		},
		OrderEmail: user.OrderEmail{
			Email: ord.Email.AsOrder(),
		},
	}
}
