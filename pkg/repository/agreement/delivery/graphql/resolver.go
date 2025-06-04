package graphql

import (
	"context"
	"fmt"
	"time"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/server/graphql/types"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/agreement"
	gqmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc *agreement.UsecaseImpl
}

func NewQueryResolver(repo agreement.Repository) *QueryResolver {
	return &QueryResolver{
		uc: agreement.NewUsecase(repo),
	}
}

// AcceptAgreement is the resolver for the acceptAgreement field.
func (r *QueryResolver) Accept(ctx context.Context, codename string, date types.DateTime, signature string) (*gqmodels.Agreement, error) {
	agreement, err := r.uc.Accept(ctx, codename, signature)
	if err != nil {
		return nil, fmt.Errorf("failed to accept agreement: %w", err)
	}
	return convertAgreement(agreement), nil
}

// Agreement is the resolver for the agreement field.
func (r *QueryResolver) Get(ctx context.Context, codename string) (*gqmodels.Agreement, error) {
	agreement, err := r.uc.Get(ctx, codename)
	if err != nil {
		return nil, fmt.Errorf("failed to get agreement: %w", err)
	}
	return convertAgreement(agreement), nil
}

// ListAgreements is the resolver for the lsitAgreements field.
func (r *QueryResolver) List(ctx context.Context, accepted *bool) ([]*gqmodels.Agreement, error) {
	agreements, err := r.uc.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list agreements: %w", err)
	}
	return xtypes.SliceApply(agreements, convertAgreement), nil
}

// NextAgreement is the resolver for the nextAgreement field.
func (r *QueryResolver) NextAwailable(ctx context.Context) (*gqmodels.Agreement, error) {
	agreement, err := r.uc.NextAwailable(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get next available agreement: %w", err)
	}
	return convertAgreement(agreement), nil
}

func convertAgreement(agreement *models.Agreement) *gqmodels.Agreement {
	if agreement == nil {
		return nil
	}
	return &gqmodels.Agreement{
		Codename:        agreement.Codename,
		Version:         agreement.Version,
		Title:           agreement.Title,
		TextMd:          agreement.BodyMarkdown,
		TextHTML:        agreement.BodyHTML,
		Type:            convertType(agreement.Type),
		AcceptAccountID: agreement.AcceptAccountID,
		AcceptByUserID:  agreement.AcceptByUserID,
		Signature:       strPtr(agreement.Signature),
		AcceptedAt:      datePtr(agreement.AcceptedAt),
		CreatedAt:       types.DateTime(agreement.CreatedAt),
	}
}

func convertType(agreementType string) gqmodels.AgreementType {
	switch agreementType {
	case "license":
		return gqmodels.AgreementTypeLicense
	case "terms_of_use":
		return gqmodels.AgreementTypeTermsOfUse
	case "contract":
		return gqmodels.AgreementTypeContract
	default:
		return gqmodels.AgreementTypeUnknown
	}
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func datePtr(t *time.Time) *types.DateTime {
	if t == nil || t.IsZero() {
		return nil
	}
	nt := types.DateTimeFromPtr(t)
	return &nt
}
