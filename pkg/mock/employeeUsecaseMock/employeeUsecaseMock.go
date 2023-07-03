// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecases/interfaces/employeeInterface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	domain "shiftsync/pkg/domain"
	request "shiftsync/pkg/helper/request"
	response "shiftsync/pkg/helper/response"

	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeUseCase is a mock of EmployeeUseCase interface.
type MockEmployeeUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeUseCaseMockRecorder
}

// MockEmployeeUseCaseMockRecorder is the mock recorder for MockEmployeeUseCase.
type MockEmployeeUseCaseMockRecorder struct {
	mock *MockEmployeeUseCase
}

// NewMockEmployeeUseCase creates a new mock instance.
func NewMockEmployeeUseCase(ctrl *gomock.Controller) *MockEmployeeUseCase {
	mock := &MockEmployeeUseCase{ctrl: ctrl}
	mock.recorder = &MockEmployeeUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeUseCase) EXPECT() *MockEmployeeUseCaseMockRecorder {
	return m.recorder
}

// AddForm mocks base method.
func (m *MockEmployeeUseCase) AddForm(r context.Context, form domain.Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddForm", r, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddForm indicates an expected call of AddForm.
func (mr *MockEmployeeUseCaseMockRecorder) AddForm(r, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddForm", reflect.TypeOf((*MockEmployeeUseCase)(nil).AddForm), r, form)
}

// ApplyLeave mocks base method.
func (m *MockEmployeeUseCase) ApplyLeave(ctx context.Context, leave domain.Leave) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyLeave", ctx, leave)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyLeave indicates an expected call of ApplyLeave.
func (mr *MockEmployeeUseCaseMockRecorder) ApplyLeave(ctx, leave interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLeave", reflect.TypeOf((*MockEmployeeUseCase)(nil).ApplyLeave), ctx, leave)
}

// Attendance mocks base method.
func (m *MockEmployeeUseCase) Attendance(ctx context.Context, id int) ([]response.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attendance", ctx, id)
	ret0, _ := ret[0].([]response.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attendance indicates an expected call of Attendance.
func (mr *MockEmployeeUseCaseMockRecorder) Attendance(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attendance", reflect.TypeOf((*MockEmployeeUseCase)(nil).Attendance), ctx, id)
}

// FormStatus mocks base method.
func (m *MockEmployeeUseCase) FormStatus(ctx context.Context, empID int) (response.FormStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormStatus", ctx, empID)
	ret0, _ := ret[0].(response.FormStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormStatus indicates an expected call of FormStatus.
func (mr *MockEmployeeUseCaseMockRecorder) FormStatus(ctx, empID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormStatus", reflect.TypeOf((*MockEmployeeUseCase)(nil).FormStatus), ctx, empID)
}

// GetDataForSalarySlip mocks base method.
func (m *MockEmployeeUseCase) GetDataForSalarySlip(ctx context.Context, id int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataForSalarySlip", ctx, id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataForSalarySlip indicates an expected call of GetDataForSalarySlip.
func (mr *MockEmployeeUseCaseMockRecorder) GetDataForSalarySlip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataForSalarySlip", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetDataForSalarySlip), ctx, id)
}

// GetDutySchedules mocks base method.
func (m *MockEmployeeUseCase) GetDutySchedules(ctx context.Context, id int) (response.Duty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDutySchedules", ctx, id)
	ret0, _ := ret[0].(response.Duty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDutySchedules indicates an expected call of GetDutySchedules.
func (mr *MockEmployeeUseCaseMockRecorder) GetDutySchedules(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDutySchedules", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetDutySchedules), ctx, id)
}

// GetLeaveStatusHistory mocks base method.
func (m *MockEmployeeUseCase) GetLeaveStatusHistory(ctx context.Context, id int) ([]response.LeaveHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeaveStatusHistory", ctx, id)
	ret0, _ := ret[0].([]response.LeaveHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeaveStatusHistory indicates an expected call of GetLeaveStatusHistory.
func (mr *MockEmployeeUseCaseMockRecorder) GetLeaveStatusHistory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeaveStatusHistory", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetLeaveStatusHistory), ctx, id)
}

// GetSalaryDetails mocks base method.
func (m *MockEmployeeUseCase) GetSalaryDetails(ctx context.Context, id int) (response.Salarydetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSalaryDetails", ctx, id)
	ret0, _ := ret[0].(response.Salarydetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalaryDetails indicates an expected call of GetSalaryDetails.
func (mr *MockEmployeeUseCaseMockRecorder) GetSalaryDetails(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalaryDetails", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetSalaryDetails), ctx, id)
}

// GetSalaryHistory mocks base method.
func (m *MockEmployeeUseCase) GetSalaryHistory(ctx context.Context, id int) ([]response.Salaryhistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSalaryHistory", ctx, id)
	ret0, _ := ret[0].([]response.Salaryhistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalaryHistory indicates an expected call of GetSalaryHistory.
func (mr *MockEmployeeUseCaseMockRecorder) GetSalaryHistory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalaryHistory", reflect.TypeOf((*MockEmployeeUseCase)(nil).GetSalaryHistory), ctx, id)
}

// Login mocks base method.
func (m *MockEmployeeUseCase) Login(r context.Context, login domain.Employee) (domain.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", r, login)
	ret0, _ := ret[0].(domain.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockEmployeeUseCaseMockRecorder) Login(r, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockEmployeeUseCase)(nil).Login), r, login)
}

// PunchIn mocks base method.
func (m *MockEmployeeUseCase) PunchIn(ctx context.Context, ID int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchIn", ctx, ID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PunchIn indicates an expected call of PunchIn.
func (mr *MockEmployeeUseCaseMockRecorder) PunchIn(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchIn", reflect.TypeOf((*MockEmployeeUseCase)(nil).PunchIn), ctx, ID)
}

// PunchOut mocks base method.
func (m *MockEmployeeUseCase) PunchOut(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchOut", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchOut indicates an expected call of PunchOut.
func (mr *MockEmployeeUseCaseMockRecorder) PunchOut(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchOut", reflect.TypeOf((*MockEmployeeUseCase)(nil).PunchOut), ctx, id)
}

// SignUp mocks base method.
func (m *MockEmployeeUseCase) SignUp(r context.Context, signup domain.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", r, signup)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockEmployeeUseCaseMockRecorder) SignUp(r, signup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockEmployeeUseCase)(nil).SignUp), r, signup)
}

// SignUpOtp mocks base method.
func (m *MockEmployeeUseCase) SignUpOtp(r context.Context, find domain.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpOtp", r, find)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUpOtp indicates an expected call of SignUpOtp.
func (mr *MockEmployeeUseCaseMockRecorder) SignUpOtp(r, find interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpOtp", reflect.TypeOf((*MockEmployeeUseCase)(nil).SignUpOtp), r, find)
}

// VerifyOtpForPunchin mocks base method.
func (m *MockEmployeeUseCase) VerifyOtpForPunchin(ctx context.Context, id int, otp request.OTPStruct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOtpForPunchin", ctx, id, otp)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyOtpForPunchin indicates an expected call of VerifyOtpForPunchin.
func (mr *MockEmployeeUseCaseMockRecorder) VerifyOtpForPunchin(ctx, id, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOtpForPunchin", reflect.TypeOf((*MockEmployeeUseCase)(nil).VerifyOtpForPunchin), ctx, id, otp)
}