package appinit

import (
	"context"
	"strings"

	"github.com/demdxx/rbac"

	"github.com/geniusrabbit/blaze-api/model"
	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	"github.com/geniusrabbit/blaze-api/pkg/permissions"
	"github.com/geniusrabbit/blaze-api/repository/account/repository"

	"github.com/sspserver/api/models"
)

var (
	crudPermissions = []string{
		acl.PermView, acl.PermList, acl.PermCount, acl.PermUpdate, acl.PermCreate, acl.PermDelete,
	}
	crudPermissionsWithApprove = append(crudPermissions, acl.PermApprove, acl.PermReject)
)

var (
	RBACStatistic = acl.RBACType{ResourceName: "statistic"}
)

const (
	PermAccountRegister = `account.register`
	PermPermissionList  = `permission.list`
	PermUserPassReset   = `password.reset`
	PermUserPassSet     = `password.set`
)

// InitModelPermissions models
func InitModelPermissions(pm *permissions.Manager) {
	// Register permission objects
	acl.InitModelPermissions(pm,
		&models.User{},
		&models.Account{},
		&models.AccountMember{},
		&models.Role{},
		&models.AuthClient{},
		&models.AccountSocialSession{},
		&models.AccountSocial{},
		&models.HistoryAction{},
		&models.Option{},
		&models.DirectAccessToken{},
		// Current project models
		&models.RTBSource{},
		&models.Application{},
		&models.Zone{},
		&models.Browser{},
		&models.OS{},
		&models.DeviceType{},
		&models.DeviceModel{},
		&models.DeviceMaker{},
		&models.Category{},
		&RBACStatistic,
		&models.Format{},
	)

	// Register user permissions
	_ = pm.RegisterNewOwningPermissions(&models.User{}, append(crudPermissionsWithApprove, PermUserPassReset, PermUserPassSet))

	// Register basic models CRUD permissions for Account with member checks
	_ = pm.RegisterNewOwningPermissions(&models.Account{}, crudPermissionsWithApprove, rbac.WithCustomCheck(accountCustomCheck))
	_ = pm.RegisterNewPermission(nil, PermAccountRegister, rbac.WithoutCustomCheck,
		rbac.WithDescription("Register new account"))

	// Register basic permissions for the AccountMember model
	_ = pm.RegisterNewOwningPermissions(&models.AccountMember{}, crudPermissionsWithApprove)
	_ = pm.RegisterNewPermissions(&models.AccountMember{}, []string{`roles.set.account`, `roles.set.all`, `invite`})

	// Register basic roles permissions
	_ = pm.RegisterNewOwningPermissions(&models.Role{}, crudPermissions)
	_ = pm.RegisterNewPermission(&models.Role{}, `check`,
		rbac.WithDescription("Check role permissions is assigned to the user"))
	_ = pm.RegisterNewPermission(nil, PermPermissionList, rbac.WithDescription("List all permissions"))

	// Register basic permissions for the AuthClient model
	_ = pm.RegisterNewOwningPermissions(&models.AuthClient{}, crudPermissions)

	// Register basic permissions for the AccountSocial model
	_ = pm.RegisterNewOwningPermissions(&models.AccountSocial{}, append(crudPermissions, `disconnect`))

	// Register basic permissions for the HistoryAction model
	_ = pm.RegisterNewOwningPermissions(&models.HistoryAction{}, []string{acl.PermView, acl.PermList, acl.PermCount})

	// Register basic permissions for the Option model
	_ = pm.RegisterNewOwningPermissions(&models.Option{}, []string{acl.PermGet, acl.PermSet, acl.PermList, acl.PermCount})

	// Register basic permissions for the DirectAccessToken model
	_ = pm.RegisterNewOwningPermissions(&models.DirectAccessToken{}, []string{acl.PermGet, acl.PermList, acl.PermCount, acl.PermCreate, acl.PermDelete})

	// =========== Current project models ===========
	// Register basic permissions for the RTBSource model
	_ = pm.RegisterNewOwningPermissions(&models.RTBSource{}, crudPermissionsWithApprove,
		rbac.WithDescription("RTB Source model permissions"))

	// Register basic permissions for the Application model
	_ = pm.RegisterNewOwningPermissions(&models.Application{}, crudPermissionsWithApprove,
		rbac.WithDescription("Application model permissions"))

	// Register basic permissions for the Zone model
	_ = pm.RegisterNewOwningPermissions(&models.Zone{}, crudPermissionsWithApprove,
		rbac.WithDescription("Zone model permissions"))

	// Register basic permissions for the Browser model
	_ = pm.RegisterNewOwningPermissions(&models.Browser{}, crudPermissions,
		rbac.WithDescription("Browser model permissions"))

	// Register basic permissions for the OS model
	_ = pm.RegisterNewOwningPermissions(&models.OS{}, crudPermissions,
		rbac.WithDescription("OS model permissions"))

	// Register basic permissions for the DeviceType model
	_ = pm.RegisterNewOwningPermissions(&models.DeviceType{}, crudPermissions,
		rbac.WithDescription("Device Type model permissions"))

	// Register basic permissions for the DeviceModel model
	_ = pm.RegisterNewOwningPermissions(&models.DeviceModel{}, crudPermissions,
		rbac.WithDescription("Device Model model permissions"))

	// Register basic permissions for the DeviceMaker model
	_ = pm.RegisterNewOwningPermissions(&models.DeviceMaker{}, crudPermissions,
		rbac.WithDescription("Device Maker model permissions"))

	// Register basic permissions for the Category model
	_ = pm.RegisterNewOwningPermissions(&models.Category{}, crudPermissions,
		rbac.WithDescription("Category model permissions"))

	// Register basic permissions for the Statistic model
	_ = pm.RegisterNewOwningPermissions(&RBACStatistic, []string{acl.PermView, acl.PermList, acl.PermCount},
		rbac.WithDescription("Statistic model permissions"))

	// Register basic permissions for the Format model
	_ = pm.RegisterNewOwningPermissions(&models.Format{}, crudPermissions,
		rbac.WithDescription("Format model permissions"))

	// =========== Register default roles ===========
	pm.RegisterRole(context.Background(),
		// Register anonymous role and fill permissions for it
		rbac.MustNewRole(permissions.AnonymousDefaultRole,
			rbac.WithDescription("Anonymous user role"),
			rbac.WithPermissions(
				`user.{view|list|count}.owner`,
				`account.{view|list|count}.owner`, PermAccountRegister,
				`role.check`,
			),
		),
		// Register default role and fill permissions for it
		rbac.MustNewRole(permissions.DefaultRole,
			rbac.WithDescription("Default user role"),
			rbac.WithPermissions(
				`user.{view|list|count}.owner`, `user.password.{set|reset}.owner`,
				`account.{view|list|count}.owner`, PermAccountRegister,
				`option.{get|set|list|count}.owner`,
				`directaccesstoken.{view|list|count}.owner`,
				`role.check`,
				`rtb_source.{view|list|count}.owner`,
				`geo_country.{view|list|count}.owner`, `geo_continent.{view|list|count}.owner`,
				`statistic.{view|list|count}.owner`,
			),
		),
	)
}

func accountCustomCheck(ctx context.Context, resource any, perm rbac.Permission) bool {
	if strings.HasSuffix(perm.Name(), `.system`) || strings.HasSuffix(perm.Name(), `.all`) {
		return true
	}
	account, _ := resource.(*model.Account)
	user := session.User(ctx)
	if account.IsOwnerUser(user.ID) {
		return true
	}
	if account.ID > 0 {
		repo := repository.New()
		if perm.MatchPermissionPattern(`*.{view|list|count}.*`) {
			return repo.IsMember(ctx, user.ID, account.ID)
		}
		return repo.IsAdmin(ctx, user.ID, account.ID)
	}
	return false
}
