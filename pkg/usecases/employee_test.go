package usecases

import (
	"context"
	"errors"
	"shiftsync/pkg/domain"
	mock "shiftsync/pkg/mock/employeeRepoMock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	mockRepo := mock.NewMockEmployeeRepository(cntrl)

	employeeUsecase := NewEmployeeUseCase(mockRepo)

	testData := []struct {
		responce    string
		employee    domain.Employee
		expectedErr error
		beforeTest  func(employeeRepo *mock.MockEmployeeRepository)
	}{
		{
			responce: "Test success",
			employee: domain.Employee{
				First_name: "Ashiq",
				Last_name:  "Sabith",
				Phone:      8606863748,
				Email:      "ashiqsabith328@gmail.com",
				User_name:  "ashiqsabith328",
				Pass_word:  "Ashiq@123",
			},
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().AddEmployee(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedErr: nil,
		},
		{
			responce: "Test failure",
			employee: domain.Employee{
				First_name: "Ashiq",
				Last_name:  "Sabith",
				Phone:      8606863748,
				Email:      "ashiqsabith328@gmail.com",
				User_name:  "ashiqsabith328",
				Pass_word:  "Ashiq@123",
			},
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().AddEmployee(gomock.Any(), gomock.Any()).Return(errors.New("Employee already exists"))
			},
			expectedErr: errors.New("Employee already exists"),
		},
	}

	for _, testCase := range testData {
		t.Run(testCase.responce, func(t *testing.T) {
			testCase.beforeTest(mockRepo)
			err := employeeUsecase.SignUp(context.Background(), testCase.employee)
			assert.Equal(t, testCase.expectedErr, err)

		})
	}
}
