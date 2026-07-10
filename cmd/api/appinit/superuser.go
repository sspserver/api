package appinit

import (
	"context"

	"go.uber.org/zap"

	blazeacl "github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/ctxlogger"
	pkgModels "github.com/geniusrabbit/blaze-api/pkg/models"
	"github.com/geniusrabbit/blaze-api/repository/account"
	accountModels "github.com/geniusrabbit/blaze-api/repository/account/models"
	userModels "github.com/geniusrabbit/blaze-api/repository/user/models"

	"github.com/sspserver/api/pkg/models"
	userstack "github.com/sspserver/api/pkg/user"
)

// EnsureSuperuser creates the superuser and system account if they do not already exist.
// Idempotent: returns nil immediately when a user with the given email is found.
func EnsureSuperuser(
	ctx context.Context,
	email, password string,
	userRepo userstack.Repository[*models.User],
	accountRepo account.SessionRepository[*models.User, *models.Account],
	memberRepo account.MemberRepository[*models.User, *models.Account],
) error {
	if email == "" || password == "" || userRepo == nil || accountRepo == nil || memberRepo == nil {
		return nil
	}

	// Bypass ACL checks. This runs at init time with no authenticated session.
	ctx = blazeacl.WithNoPermCheck(ctx)

	lg := ctxlogger.Get(ctx)
	lg.Info("Ensuring superuser exists", zap.String("email", email))

	// Check whether the user already exists.
	existing, _ := userRepo.GetByEmail(ctx, email)
	if existing != nil && existing.GetID() != 0 {
		lg.Info("Superuser already exists, skipping creation", zap.String("email", email))
		return nil
	}

	// Create superuser.
	u := &models.User{
		UserBase:  userModels.UserBase{Approve: pkgModels.ApprovedApproveStatus},
		UserEmail: userModels.UserEmail{Email: email},
	}
	userID, err := userRepo.CreateWithPassword(ctx, u, password)
	if err != nil {
		lg.Error("Failed to create superuser", zap.String("email", email), zap.Error(err))
		return err
	}
	u.ID = userID

	// Create system account and link superuser as admin member.
	acc := &models.Account{
		AccountBase:      accountModels.AccountBase{Approve: pkgModels.ApprovedApproveStatus},
		Name:             "system",
		Description:      "System account",
		CountryCode:      "",
		City:             "",
		ZipCode:          "",
		Address:          "",
		Phone:            "",
		VATNumber:        "",
		CompanyRegNumber: "",
	}
	accID, err := accountRepo.Create(ctx, acc)
	if err != nil {
		lg.Error("Failed to create system account", zap.String("email", email), zap.Error(err))
		return err
	}
	acc.ID = accID

	if err = memberRepo.LinkMember(ctx, acc, true, u); err != nil {
		lg.Error("Failed to link user as admin member", zap.String("email", email), zap.Error(err))
		return err
	}

	if err = memberRepo.SetMemberRoles(ctx, acc, u, "system:admin"); err != nil {
		lg.Error("Failed to assign system:admin role", zap.String("email", email), zap.Error(err))
		return err
	}

	lg.Info("Superuser and system account created successfully",
		zap.String("email", email), zap.Uint64("user_id", userID), zap.Uint64("account_id", accID))

	return nil
}
