package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
)

type AdminUseCase interface {
	SignUp(ctx context.Context, admin domain.Admin) error
	SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error)
	Applications(ctx context.Context) ([]response.Form, error)
	ApproveApplication(ctx context.Context, form domain.Form) error
	FormCorrection(ctx context.Context, form domain.Form) error
}
