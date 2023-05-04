package interfaces

import (
	"context"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, find domain.Admin) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) error
	GetAllForms(ctx context.Context) ([]response.Form, error)
	ApproveApplication(ctx context.Context, form domain.Form, id int)
	FindFormByID(ctx context.Context, fID int) error
	FormCorrection(ctx context.Context, form domain.Form)
}
