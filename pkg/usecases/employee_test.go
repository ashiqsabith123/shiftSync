package usecases

import (
	"shiftsync/pkg/domain"
	mock "shiftsync/pkg/mock/employeeRepoMock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSignup(t *testing.T) {

	cntrl := gomock.NewController(t)

	mockRepo := mock.NewMockEmployeeRepository(cntrl)

	employeeUsecase := NewEmployeeUseCase(mockRepo)

	testData := []struct {
		name        string
		employee    domain.Employee
		expectedErr error
	}{}

}
