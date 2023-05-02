package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, find domain.Admin) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	GetAllForms(ctx context.Context) ([]domain.Form, error)
	ApproveApplication(ctx context.Context, form domain.Form)
	FindFormByID(ctx context.Context, fID int) error
	FormCorrection(ctx context.Context, form domain.Form)
}
