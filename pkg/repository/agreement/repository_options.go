package agreement

import (
	"context"
	"time"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	"github.com/geniusrabbit/blaze-api/repository/option"

	"github.com/sspserver/api/pkg/models"
)

type acceptance struct {
	AccountID  uint64    `json:"account_id"`
	UserID     uint64    `json:"user_id"`
	AcceptedAt time.Time `json:"accepted_at"`
	Signature  string    `json:"signature,omitempty"`
}

type RepositoryOptionsImpl struct {
	opts       option.Usecase
	agreements []*models.Agreement
}

func NewRepositoryOptions(opts option.Usecase, agreements []*models.Agreement) *RepositoryOptionsImpl {
	return &RepositoryOptionsImpl{
		opts:       opts,
		agreements: agreements,
	}
}

// Get retrieves an agreement by its codename and fills acceptance details for the current user and account.
func (r *RepositoryOptionsImpl) Get(ctx context.Context, codename string) (*models.Agreement, error) {
	for _, agreement := range r.agreements {
		if agreement.Codename == codename {
			// Check if the agreement is accepted
			r.fillAgreementAcceptance(ctx, agreement)
			return agreement, nil
		}
	}
	return nil, ErrAgreementNotFound
}

// List returns all agreements with acceptance details for the current user and account.
func (r *RepositoryOptionsImpl) List(ctx context.Context) ([]*models.Agreement, error) {
	newList := make([]*models.Agreement, 0, len(r.agreements))
	for _, agreement := range r.agreements {
		newAgreement := *agreement
		r.fillAgreementAcceptance(ctx, &newAgreement)
		newList = append(newList, &newAgreement)
	}
	return newList, nil
}

// NextAwailable returns the next available agreement for the current user and account.
func (r *RepositoryOptionsImpl) NextAwailable(ctx context.Context) (*models.Agreement, error) {
	for _, agreement := range r.agreements {
		newAgreement := *agreement
		if !r.fillAgreementAcceptance(ctx, &newAgreement) {
			return &newAgreement, nil
		}
	}
	return nil, nil
}

// Accept accepts an agreement for the current user and account.
func (r *RepositoryOptionsImpl) Accept(ctx context.Context, codename string, signature string) (*models.Agreement, error) {
	agreement, err := r.Get(ctx, codename) // Ensure the agreement exists
	if err != nil {
		return nil, err
	}

	// Check if the agreement is already accepted
	if agreement.AcceptAccountID != 0 && agreement.AcceptedAt != nil && !agreement.AcceptedAt.IsZero() {
		return nil, ErrAgreementAlreadyAccepted
	}

	// Create acceptance record
	accept := acceptance{
		AccountID:  session.Account(ctx).ID,
		UserID:     session.User(ctx).ID,
		AcceptedAt: time.Now(),
		Signature:  signature,
	}

	err = r.opts.SetOption(ctx,
		"agreement_"+agreement.Codename,
		models.AccountOptionType,
		session.Account(ctx).ID,
		accept,
	)
	if err != nil {
		return nil, err
	}

	// Update the agreement with acceptance details
	agreement.AcceptAccountID = accept.AccountID
	agreement.AcceptByUserID = accept.UserID
	agreement.AcceptedAt = &accept.AcceptedAt
	agreement.Signature = accept.Signature

	return agreement, nil
}

func (r *RepositoryOptionsImpl) fillAgreementAcceptance(ctx context.Context, agreement *models.Agreement) bool {
	sessAccountID := session.Account(ctx).ID
	opt, err := r.opts.Get(ctx, "agreement_"+agreement.Codename, models.AccountOptionType, sessAccountID)
	if err != nil || opt == nil || opt.Value.Data == nil {
		return false
	}

	var acc acceptance
	if err := gocast.TryCopyStructContext(ctx, &acc, opt.Value.Data, `json`); err != nil {
		return false
	}

	if acc.AccountID != sessAccountID || acc.AcceptedAt.IsZero() || acc.AcceptedAt.After(time.Now()) {
		return false
	}

	agreement.AcceptAccountID = acc.AccountID
	agreement.AcceptByUserID = acc.UserID
	agreement.AcceptedAt = new(time.Time)
	*agreement.AcceptedAt = acc.AcceptedAt
	agreement.Signature = acc.Signature

	return true
}
