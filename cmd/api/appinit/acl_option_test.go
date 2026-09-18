package appinit

import (
	"context"
	"testing"

	"github.com/demdxx/rbac"

	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	"github.com/geniusrabbit/blaze-api/pkg/permissions"

	"github.com/sspserver/api/pkg/models"
)

const (
	testUserID    = uint64(10)
	testAccountID = uint64(20)
)

func TestOptionACL(t *testing.T) {
	t.Parallel()

	blank := &models.Option{}
	ownUser := &models.Option{Type: models.UserOptionType, TargetID: testUserID}
	otherUser := &models.Option{Type: models.UserOptionType, TargetID: testUserID + 1}
	ownAccount := &models.Option{Type: models.AccountOptionType, TargetID: testAccountID}
	otherAccount := &models.Option{Type: models.AccountOptionType, TargetID: testAccountID + 1}
	systemOpt := &models.Option{Type: models.SystemOptionType}

	ownerCtx, ownerAcc := optionACLSession(t, `option.{get|set|list|count}.owner`)
	adminCtx, adminAcc := optionACLSession(t, `option.{get|set|list|count}.{account|owner}`)
	allCtx, allAcc := optionACLSession(t, `option.{get|set|list|count}.all`)

	tests := []struct {
		name    string
		ctx     context.Context
		account *models.Account
		opt     *models.Option
		want    bool
	}{
		{name: "blank GraphQL option with owner cover", ctx: ownerCtx, account: ownerAcc, opt: blank, want: true},
		{name: "USER own id with owner cover", ctx: ownerCtx, account: ownerAcc, opt: ownUser, want: true},
		{name: "USER foreign id with owner cover", ctx: ownerCtx, account: ownerAcc, opt: otherUser, want: false},
		{name: "ACCOUNT own id with owner cover", ctx: ownerCtx, account: ownerAcc, opt: ownAccount, want: false},
		{name: "ACCOUNT foreign id with owner cover", ctx: ownerCtx, account: ownerAcc, opt: otherAccount, want: false},
		{name: "SYSTEM with owner cover", ctx: ownerCtx, account: ownerAcc, opt: systemOpt, want: false},

		{name: "blank GraphQL option with account cover", ctx: adminCtx, account: adminAcc, opt: blank, want: true},
		{name: "USER own id with account+owner cover", ctx: adminCtx, account: adminAcc, opt: ownUser, want: true},
		{name: "ACCOUNT own id with account cover", ctx: adminCtx, account: adminAcc, opt: ownAccount, want: true},
		{name: "ACCOUNT foreign id with account cover", ctx: adminCtx, account: adminAcc, opt: otherAccount, want: false},
		{name: "SYSTEM with account cover", ctx: adminCtx, account: adminAcc, opt: systemOpt, want: false},

		{name: "SYSTEM with all cover", ctx: allCtx, account: allAcc, opt: systemOpt, want: true},
		{name: "ACCOUNT with all cover", ctx: allCtx, account: allAcc, opt: ownAccount, want: true},
		{name: "USER with all cover", ctx: allCtx, account: allAcc, opt: ownUser, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.account.CheckPermissions(tt.ctx, tt.opt, `option.set.*`)
			if got != tt.want {
				t.Fatalf("CheckPermissions(option.set.*) = %v, want %v", got, tt.want)
			}
		})
	}
}

func optionACLSession(t *testing.T, permPattern string) (context.Context, *models.Account) {
	t.Helper()

	pm := &permissions.Manager{Manager: rbac.NewManager(nil)}
	InitModelPermissions(pm)

	ctx := permissions.WithManager(context.Background(), pm)
	roleName := "opt-test:" + permPattern
	pm.RegisterRole(ctx, rbac.MustNewRole(roleName, rbac.WithPermissions(permPattern)))

	user := &models.User{}
	user.SetID(testUserID)

	acc := &models.Account{}
	acc.SetID(testAccountID)
	acc.SetPermissions(pm.Role(ctx, roleName))

	return session.WithUserAccount(ctx, user, acc), acc
}
