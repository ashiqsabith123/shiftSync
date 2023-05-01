package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type AdminUseCase interface {
	SignUp(ctx context.Context, admin domain.Admin) error
	SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error)
	Applications(ctx context.Context) ([]domain.Form, error)
}
