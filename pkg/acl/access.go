package acl

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
)

// WithNoPermCheck returns new context with disabled permission check
func WithNoPermCheck(ctx context.Context) context.Context {
	return acl.WithNoPermCheck(ctx)
}

// IsNoPermCheck returns `true` if the permission check is disabled
func IsNoPermCheck(ctx context.Context) bool {
	return acl.IsNoPermCheck(ctx)
}

// The permission list
const (
	PermView      = acl.PermView
	PermCreate    = acl.PermCreate
	PermUpdate    = acl.PermUpdate
	PermDelete    = acl.PermDelete
	PermList      = acl.PermList
	PermAuthCross = acl.PermAuthCross
	PermCount     = acl.PermCount
	PermApprove   = acl.PermApprove
	PermReject    = acl.PermReject
	PermGet       = acl.PermGet
	PermSet       = acl.PermSet
	PermRun       = `run`
	PermPause     = `pause`
)

// HavePermissions returns `true` if the `user` have all permissions from the list
func HavePermissions(ctx context.Context, permissions ...string) bool {
	return acl.HavePermissions(ctx, permissions...)
}

// HaveObjectPermissions returns `true` if the `user` have all permissions from the list for the object
func HaveObjectPermissions(ctx context.Context, obj any, permissions ...string) bool {
	return acl.HaveObjectPermissions(ctx, obj, permissions...)
}

// HaveAccessView to the object returns `true` if user can read of the object
func HaveAccessView(ctx context.Context, obj any) bool {
	return acl.HaveAccessView(ctx, obj)
}

// HaveAccessList to the object returns `true` if user can read list of the object
func HaveAccessList(ctx context.Context, obj any) bool {
	return acl.HaveAccessList(ctx, obj)
}

// HaveAccessCount of the object returns `true` if user can count the object
func HaveAccessCount(ctx context.Context, obj any) bool {
	return acl.HaveAccessCount(ctx, obj)
}

// HaveAccessCreate of the object returns `true` if user can create this type of object
func HaveAccessCreate(ctx context.Context, obj any) bool {
	return acl.HaveAccessCreate(ctx, obj)
}

// HaveAccessUpdate of the object returns `true` if user can update the object
func HaveAccessUpdate(ctx context.Context, obj any) bool {
	return acl.HaveAccessUpdate(ctx, obj)
}

// HaveAccessDelete of the object returns `true` if user can delite the object
func HaveAccessDelete(ctx context.Context, obj any) bool {
	return acl.HaveAccessDelete(ctx, obj)
}

// HaveAccessApprove of the object returns `true` if user can approve the object
func HaveAccessApprove(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermApprove+`.*`)
}

// HaveAccessReject of the object returns `true` if user can reject the object
func HaveAccessReject(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermReject+`.*`)
}

// HaveAccessRun of the object returns `true` if user can run the object
func HaveAccessRun(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermRun+`.*`)
}

// HaveAccessPause of the object returns `true` if user can pause the object
func HaveAccessPause(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermPause+`.*`)
}

// HaveAccessGet of the object returns `true` if user can get the object
func HaveAccessGet(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermGet+`.*`)
}

// HaveAccessSet of the object returns `true` if user can set the object
func HaveAccessSet(ctx context.Context, obj any) bool {
	return acl.HaveObjectPermissions(ctx, obj, PermSet+`.*`)
}

// HaveAccountLink of the object to the current account
func HaveAccountLink(ctx context.Context, obj any) bool {
	return acl.HaveAccountLink(ctx, obj)
}

// HasPermission returns `true` if the `user` have all permissions from the list (without custom check)
func HasPermission(ctx context.Context, permissions ...string) bool {
	return acl.HasPermission(ctx, permissions...)
}
