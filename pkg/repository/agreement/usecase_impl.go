package agreement

import (
	"context"

	"github.com/sspserver/api/pkg/models"
)

type UsecaseImpl struct {
	repo Repository
}

func NewUsecase(repo Repository) *UsecaseImpl {
	return &UsecaseImpl{
		repo: repo,
	}
}

func (u *UsecaseImpl) Get(ctx context.Context, codename string) (*models.Agreement, error) {
	agreement, err := u.repo.Get(ctx, codename)
	if err != nil {
		return nil, err
	}
	return agreement, nil
}

func (u *UsecaseImpl) List(ctx context.Context) ([]*models.Agreement, error) {
	agreements, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return agreements, nil
}

func (u *UsecaseImpl) NextAwailable(ctx context.Context) (*models.Agreement, error) {
	agreement, err := u.repo.NextAwailable(ctx)
	if err != nil {
		return nil, err
	}
	return agreement, nil
}

func (u *UsecaseImpl) Accept(ctx context.Context, codename string, signature string) (*models.Agreement, error) {
	agreement, err := u.repo.Accept(ctx, codename, signature)
	if err != nil {
		return nil, err
	}
	return agreement, nil
}
