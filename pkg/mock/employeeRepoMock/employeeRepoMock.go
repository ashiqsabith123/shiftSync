// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/employeeinterfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	domain "shiftsync/pkg/domain"
	response "shiftsync/pkg/helper/response"

	gomock "github.com/golang/mock/gomock"
)

// MockEmployeeRepository is a mock of EmployeeRepository interface.
type MockEmployeeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeRepositoryMockRecorder
}

// MockEmployeeRepositoryMockRecorder is the mock recorder for MockEmployeeRepository.
type MockEmployeeRepositoryMockRecorder struct {
	mock *MockEmployeeRepository
}

// NewMockEmployeeRepository creates a new mock instance.
func NewMockEmployeeRepository(ctrl *gomock.Controller) *MockEmployeeRepository {
	mock := &MockEmployeeRepository{ctrl: ctrl}
	mock.recorder = &MockEmployeeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeRepository) EXPECT() *MockEmployeeRepositoryMockRecorder {
	return m.recorder
}

// AddEmployee mocks base method.
func (m *MockEmployeeRepository) AddEmployee(cntxt context.Context, signup domain.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEmployee", cntxt, signup)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEmployee indicates an expected call of AddEmployee.
func (mr *MockEmployeeRepositoryMockRecorder) AddEmployee(cntxt, signup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEmployee", reflect.TypeOf((*MockEmployeeRepository)(nil).AddEmployee), cntxt, signup)
}

// AddForm mocks base method.
func (m *MockEmployeeRepository) AddForm(cntxt context.Context, form domain.Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddForm", cntxt, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddForm indicates an expected call of AddForm.
func (mr *MockEmployeeRepositoryMockRecorder) AddForm(cntxt, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddForm", reflect.TypeOf((*MockEmployeeRepository)(nil).AddForm), cntxt, form)
}

// ApplyLeave mocks base method.
func (m *MockEmployeeRepository) ApplyLeave(ctx context.Context, leave domain.Leave) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyLeave", ctx, leave)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyLeave indicates an expected call of ApplyLeave.
func (mr *MockEmployeeRepositoryMockRecorder) ApplyLeave(ctx, leave interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyLeave", reflect.TypeOf((*MockEmployeeRepository)(nil).ApplyLeave), ctx, leave)
}

// Attendance mocks base method.
func (m *MockEmployeeRepository) Attendance(ctx context.Context, id int) ([]response.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attendance", ctx, id)
	ret0, _ := ret[0].([]response.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Attendance indicates an expected call of Attendance.
func (mr *MockEmployeeRepositoryMockRecorder) Attendance(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attendance", reflect.TypeOf((*MockEmployeeRepository)(nil).Attendance), ctx, id)
}

// CheckFormDetails mocks base method.
func (m *MockEmployeeRepository) CheckFormDetails(cntxt context.Context, form domain.Form) (domain.Form, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFormDetails", cntxt, form)
	ret0, _ := ret[0].(domain.Form)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckFormDetails indicates an expected call of CheckFormDetails.
func (mr *MockEmployeeRepositoryMockRecorder) CheckFormDetails(cntxt, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFormDetails", reflect.TypeOf((*MockEmployeeRepository)(nil).CheckFormDetails), cntxt, form)
}

// FindEmployee mocks base method.
func (m *MockEmployeeRepository) FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmployee", cntxt, find)
	ret0, _ := ret[0].(domain.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEmployee indicates an expected call of FindEmployee.
func (mr *MockEmployeeRepositoryMockRecorder) FindEmployee(cntxt, find interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmployee", reflect.TypeOf((*MockEmployeeRepository)(nil).FindEmployee), cntxt, find)
}

// FormCorrection mocks base method.
func (m *MockEmployeeRepository) FormCorrection(ctx context.Context, form domain.Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormCorrection", ctx, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// FormCorrection indicates an expected call of FormCorrection.
func (mr *MockEmployeeRepositoryMockRecorder) FormCorrection(ctx, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormCorrection", reflect.TypeOf((*MockEmployeeRepository)(nil).FormCorrection), ctx, form)
}

// FormStatus mocks base method.
func (m *MockEmployeeRepository) FormStatus(ctx context.Context, empID int) (response.FormStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormStatus", ctx, empID)
	ret0, _ := ret[0].(response.FormStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormStatus indicates an expected call of FormStatus.
func (mr *MockEmployeeRepositoryMockRecorder) FormStatus(ctx, empID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormStatus", reflect.TypeOf((*MockEmployeeRepository)(nil).FormStatus), ctx, empID)
}

// GetCountOfLeaveTaken mocks base method.
func (m *MockEmployeeRepository) GetCountOfLeaveTaken(ctx context.Context, reqCount response.LeaveCount) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountOfLeaveTaken", ctx, reqCount)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountOfLeaveTaken indicates an expected call of GetCountOfLeaveTaken.
func (mr *MockEmployeeRepositoryMockRecorder) GetCountOfLeaveTaken(ctx, reqCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountOfLeaveTaken", reflect.TypeOf((*MockEmployeeRepository)(nil).GetCountOfLeaveTaken), ctx, reqCount)
}

// GetDataForSalarySlip mocks base method.
func (m *MockEmployeeRepository) GetDataForSalarySlip(ctx context.Context, id int) (response.SalarySlip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataForSalarySlip", ctx, id)
	ret0, _ := ret[0].(response.SalarySlip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataForSalarySlip indicates an expected call of GetDataForSalarySlip.
func (mr *MockEmployeeRepositoryMockRecorder) GetDataForSalarySlip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataForSalarySlip", reflect.TypeOf((*MockEmployeeRepository)(nil).GetDataForSalarySlip), ctx, id)
}

// GetDuty mocks base method.
func (m *MockEmployeeRepository) GetDuty(ctx context.Context, id int) (response.Duty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDuty", ctx, id)
	ret0, _ := ret[0].(response.Duty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDuty indicates an expected call of GetDuty.
func (mr *MockEmployeeRepositoryMockRecorder) GetDuty(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDuty", reflect.TypeOf((*MockEmployeeRepository)(nil).GetDuty), ctx, id)
}

// GetDutySchedules mocks base method.
func (m *MockEmployeeRepository) GetDutySchedules(ctx context.Context, id int) (response.Duty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDutySchedules", ctx, id)
	ret0, _ := ret[0].(response.Duty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDutySchedules indicates an expected call of GetDutySchedules.
func (mr *MockEmployeeRepositoryMockRecorder) GetDutySchedules(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDutySchedules", reflect.TypeOf((*MockEmployeeRepository)(nil).GetDutySchedules), ctx, id)
}

// GetLastAppliedLeave mocks base method.
func (m *MockEmployeeRepository) GetLastAppliedLeave(ctx context.Context, check domain.Leave) (response.LeaveAppiled, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAppliedLeave", ctx, check)
	ret0, _ := ret[0].(response.LeaveAppiled)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastAppliedLeave indicates an expected call of GetLastAppliedLeave.
func (mr *MockEmployeeRepositoryMockRecorder) GetLastAppliedLeave(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAppliedLeave", reflect.TypeOf((*MockEmployeeRepository)(nil).GetLastAppliedLeave), ctx, check)
}

// GetSalaryDetails mocks base method.
func (m *MockEmployeeRepository) GetSalaryDetails(ctx context.Context, id int) (response.Salarydetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSalaryDetails", ctx, id)
	ret0, _ := ret[0].(response.Salarydetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalaryDetails indicates an expected call of GetSalaryDetails.
func (mr *MockEmployeeRepositoryMockRecorder) GetSalaryDetails(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalaryDetails", reflect.TypeOf((*MockEmployeeRepository)(nil).GetSalaryDetails), ctx, id)
}

// GetSalaryHistory mocks base method.
func (m *MockEmployeeRepository) GetSalaryHistory(ctx context.Context, id int) ([]response.SalaryHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSalaryHistory", ctx, id)
	ret0, _ := ret[0].([]response.SalaryHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSalaryHistory indicates an expected call of GetSalaryHistory.
func (mr *MockEmployeeRepositoryMockRecorder) GetSalaryHistory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSalaryHistory", reflect.TypeOf((*MockEmployeeRepository)(nil).GetSalaryHistory), ctx, id)
}

// LeaveStatusHistory mocks base method.
func (m *MockEmployeeRepository) LeaveStatusHistory(ctx context.Context, id int) ([]response.LeaveHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveStatusHistory", ctx, id)
	ret0, _ := ret[0].([]response.LeaveHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveStatusHistory indicates an expected call of LeaveStatusHistory.
func (mr *MockEmployeeRepositoryMockRecorder) LeaveStatusHistory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveStatusHistory", reflect.TypeOf((*MockEmployeeRepository)(nil).LeaveStatusHistory), ctx, id)
}

// PunchIn mocks base method.
func (m *MockEmployeeRepository) PunchIn(ctx context.Context, punchin domain.Attendance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchIn", ctx, punchin)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchIn indicates an expected call of PunchIn.
func (mr *MockEmployeeRepositoryMockRecorder) PunchIn(ctx, punchin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchIn", reflect.TypeOf((*MockEmployeeRepository)(nil).PunchIn), ctx, punchin)
}

// PunchOut mocks base method.
func (m *MockEmployeeRepository) PunchOut(ctx context.Context, punchout domain.Attendance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PunchOut", ctx, punchout)
	ret0, _ := ret[0].(error)
	return ret0
}

// PunchOut indicates an expected call of PunchOut.
func (mr *MockEmployeeRepositoryMockRecorder) PunchOut(ctx, punchout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PunchOut", reflect.TypeOf((*MockEmployeeRepository)(nil).PunchOut), ctx, punchout)
}
