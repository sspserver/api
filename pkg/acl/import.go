package acl

import (
	"context"

	"github.com/demdxx/rbac"
	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/permissions"
)

type checkFnk func(ctx context.Context, resource any, perm rbac.Permission) bool

type RBACType = acl.RBACType

// InitModelPermissions for particular models
func InitModelPermissions(pm *permissions.Manager, models ...any) {
	acl.InitModelPermissions(pm, models...)
}

// InitModelPermissionsWithCustomCheck for particular models and extra custom check function
func InitModelPermissionsWithCustomCheck(pm *permissions.Manager, customCheck checkFnk, models ...any) {
	acl.InitModelPermissionsWithCustomCheck(pm,
		(func(ctx context.Context, resource any, perm rbac.Permission) bool)(customCheck),
		models...)
}
