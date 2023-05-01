package usecases

import (
	"context"
	"errors"
	"fmt"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/encrypt"
	"shiftsync/pkg/helper"
	repo "shiftsync/pkg/repository/interfaces"
	service "shiftsync/pkg/usecases/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepo repo.AdminRepository
}

func NewAdminUseCase(rep repo.AdminRepository) service.AdminUseCase {
	return &adminUseCase{adminRepo: rep}
}

func (a *adminUseCase) SignUp(ctx context.Context, admin domain.Admin) error {
	_, err := a.adminRepo.FindAdmin(ctx, admin)
	if err == nil {
		return errors.New("admin already exist")
	}

	hash, hasherr := bcrypt.GenerateFromPassword([]byte(admin.Pass_word), 14)

	if hasherr != nil {
		return errors.New("hashing failed" + hasherr.Error())
	}

	admin.Pass_word = string(hash)

	if err := a.adminRepo.SaveAdmin(ctx, admin); err != nil {
		return errors.New("unable to add admin " + err.Error())
	}

	return nil
}

func (a *adminUseCase) SignIn(ctx context.Context, details domain.Admin) (domain.Admin, error) {
	admin, err := a.adminRepo.FindAdmin(ctx, details)
	if err != nil {
		return details, errors.New("invalid credentials " + err.Error())
	}

	if berr := bcrypt.CompareHashAndPassword([]byte(admin.Pass_word), []byte(details.Pass_word)); berr != nil {
		return details, errors.New("incorrect password")
	}

	return admin, nil
}

func (a *adminUseCase) Applications(ctx context.Context) ([]domain.Form, error) {
	forms, err := a.adminRepo.GetAllForms(ctx)

	// for i := 0; i < len(forms); i++ {
	// 	fmt.Println("acc", forms[i].Account_no)
	// 	fmt.Println(string(helper.Decode(forms[i].Account_no)))
	// 	forms[i].Account_no = string(encrypt.Decrypt(helper.Decode(forms[i].Account_no)))
	// }

	fmt.Println("uu" + string(encrypt.Decrypt(helper.Decode(forms[0].Account_no))))

	forms[0].Account_no = string(encrypt.Decrypt(helper.Decode(forms[0].Account_no)))

	if err != nil {
		return []domain.Form{}, errors.New("no forms found")
	}

	return forms, nil

}
