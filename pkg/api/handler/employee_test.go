package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/request"
	"shiftsync/pkg/helper/response"
	mock "shiftsync/pkg/mock/employeeUsecaseMock"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func mockNeeds(t *testing.T) (*EmployeeHandler, *mock.MockEmployeeUseCase) {

	cntrl := gomock.NewController(t)
	mockUseCase := mock.NewMockEmployeeUseCase(cntrl)
	employeeHandler := NewEmployeeHandler(mockUseCase)

	return employeeHandler, mockUseCase
}

func TestPostLogin(t *testing.T) {

	employeeHandler, mockUseCase := mockNeeds(t)

	emp := domain.Employee{
		ID:        1,
		User_name: "ashiq@328",
		Pass_word: "Ashiq@123",
	}

	testData := []struct {
		name         string
		loginDetails request.LoginStruct
		response     response.Response
		beforeTest   func(employeeUseCase *mock.MockEmployeeUseCase)
	}{

		{
			name: "Test Login",
			response: response.Response{
				StatusCode: 200,
				Message:    "succesfuly logged in",
				Errors:     nil,
			},
			beforeTest: func(employeeUseCase *mock.MockEmployeeUseCase) {
				employeeUseCase.EXPECT().Login(gomock.Any(), domain.Employee{User_name: "ashiq@328", Pass_word: "Ashiq@123"}).Return(emp, nil)
			},
		},
	}

	for _, testCase := range testData {
		t.Run(testCase.name, func(t *testing.T) {

			router := gin.Default()
			router.POST("/employee/signin/", employeeHandler.PostLogin)

			data := `{"username":"ashiq@328","password":"Ashiq@123"}`

			req, _ := http.NewRequest("POST", "/employee/signin/", strings.NewReader(data))
			req.Header.Set("Content-Type", "application/json")

			respRecorder := httptest.NewRecorder()
			testCase.beforeTest(mockUseCase)

			router.ServeHTTP(respRecorder, req)

			var actual response.Response
			json.Unmarshal(respRecorder.Body.Bytes(), &actual)

			assert.Equal(t, testCase.response.StatusCode, actual.StatusCode)
			assert.Equal(t, testCase.response.Message, actual.Message)
			assert.Equal(t, testCase.response.Errors, actual.Errors)

		})
	}

}
