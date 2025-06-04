package agreement

import (
	"context"
	"errors"

	"github.com/sspserver/api/pkg/models"
)

var (
	ErrAgreementNotFound        = errors.New("agreement not found")
	ErrAgreementAlreadyAccepted = errors.New("agreement already accepted")
)

type Repository interface {
	Get(ctx context.Context, codename string) (*models.Agreement, error)
	List(ctx context.Context) ([]*models.Agreement, error)
	NextAwailable(ctx context.Context) (*models.Agreement, error)
	Accept(ctx context.Context, codename string, signature string) (*models.Agreement, error)
}
