package usecases

import (
	"context"
	"errors"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
	mock "shiftsync/pkg/mock/employeeRepoMock"
	"shiftsync/pkg/usecases/interfaces"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func mockNeeds(t *testing.T) (interfaces.EmployeeUseCase, *mock.MockEmployeeRepository) {

	cntrl := gomock.NewController(t)
	mockRepo := mock.NewMockEmployeeRepository(cntrl)
	employeeUsecase := NewEmployeeUseCase(mockRepo)

	return employeeUsecase, mockRepo
}

func TestSignup(t *testing.T) {

	employeeUsecase, mockRepo := mockNeeds(t)

	testData := []struct {
		name        string
		employee    domain.Employee
		beforeTest  func(employeeRepo *mock.MockEmployeeRepository)
		expectedErr error
	}{
		{
			name: "Signup_Accout_created_succesfully",
			employee: domain.Employee{
				First_name: "Ashiq",
				Last_name:  "Sabith",
				Phone:      8606863748,
				Email:      "ashiqsabith328@gmail.com",
				User_name:  "ashiqsabith328",
				Pass_word:  "Ashiq@123",
			},
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().AddEmployee(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testData {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.beforeTest(mockRepo)
			err := employeeUsecase.SignUp(context.Background(), testCase.employee)
			assert.Equal(t, testCase.expectedErr, err)

		})
	}
}

func TestLogin(t *testing.T) {
	employeeUsecase, mockRepo := mockNeeds(t)

	reEmp := domain.Employee{User_name: "ashiqsabith@328"}
	hash, _ := bcrypt.GenerateFromPassword([]byte("Ashiq@123"), 14)
	reEmp.Pass_word = string(hash)

	testData := []struct {
		name          string
		employee      domain.Employee
		beforeTest    func(employeeRepo *mock.MockEmployeeRepository)
		expectedError error
	}{
		{
			name: "Test success",
			employee: domain.Employee{
				User_name: "ashiqsabith@328",
				Pass_word: "Ashiq@123",
			},
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().
					FindEmployee(gomock.Any(), domain.Employee{User_name: "ashiqsabith@328", Pass_word: "Ashiq@123"}).
					Return(reEmp, nil)
			},
			expectedError: nil,
		},
		{
			name: "Test response user not exist",
			employee: domain.Employee{
				User_name: "ashiqsabith@328",
				Pass_word: "Ashiq@123",
			},
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().
					FindEmployee(gomock.Any(), domain.Employee{User_name: "ashiqsabith@328", Pass_word: "Ashiq@123"}).
					Return(domain.Employee{}, errors.New("User does not exist"))
			},
			expectedError: errors.New("User does not exist"),
		},
	}

	for _, testcase := range testData {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.beforeTest(mockRepo)
			emp, err := employeeUsecase.Login(context.Background(), testcase.employee)
			assert.Equal(t, testcase.expectedError, err)
			if err == nil {
				assert.Equal(t, reEmp, emp)
			}
		})
	}

}

func TestGetLeaveStatusHistory(t *testing.T) {
	employeeUsecase, mockRepo := mockNeeds(t)

	mockExcepthistory := []response.LeaveHistory{
		{
			Leave_type: "Casual",
			From:       "10-10-2001",
			To:         "20-10-2001",
			Status:     "A",
		},
		{
			Leave_type: "Casual",
			From:       "10-11-2001",
			To:         "20-11-2001",
			Status:     "D",
		},
	}

	actualHistory := []response.LeaveHistory{
		{
			Leave_type: "Casual",
			From:       "10-10-2001",
			To:         "20-10-2001",
			Status:     "Approved",
		},
		{
			Leave_type: "Casual",
			From:       "10-11-2001",
			To:         "20-11-2001",
			Status:     "Declined",
		},
	}

	testData := []struct {
		name        string
		id          int
		beforeTest  func(employeeRepo *mock.MockEmployeeRepository)
		expectedErr error
	}{
		{
			name: "Test Responce success",
			id:   1,
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().LeaveStatusHistory(gomock.Any(), 1).
					Return(mockExcepthistory, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Test Responce no leaves found",
			id:   1,
			beforeTest: func(employeeRepo *mock.MockEmployeeRepository) {
				employeeRepo.EXPECT().LeaveStatusHistory(gomock.Any(), 1).
					Return([]response.LeaveHistory{}, errors.New("No leave found"))
			},
			expectedErr: errors.New("No leave found"),
		},
	}

	for _, testcase := range testData {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.beforeTest(mockRepo)
			status, err := employeeUsecase.GetLeaveStatusHistory(context.Background(), testcase.id)
			if err == nil {
				assert.Equal(t, status, actualHistory)
			} else {
				assert.Equal(t, status, []response.LeaveHistory{})
			}

			assert.Equal(t, testcase.expectedErr, err)
		})
	}
}
