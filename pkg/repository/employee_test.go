package repository

import (
	"context"
	"errors"
	"log"
	"reflect"
	"shiftsync/pkg/domain"
	"shiftsync/pkg/helper/response"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetLastAppliedLeave(t *testing.T) {
	// Create a mock database connection
	mockDB, mock, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock SQL created successfully")
	}
	defer mockDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock SQL connected with GORM successfully")
	}

	// Create the employee database instance with the mock database
	employeeDB := NewEmployeeRepository(db)

	// Insert a sample leave record in the database
	leave := domain.Leave{
		From: "10-10-2001",
		To:   "20-10-2001",
	}

	// Create a sample check object
	check := domain.Leave{
		EmployeeID: 1,
	}

	// Test case 1: Get the leave dates last applied
	expectedQuery := `SELECT leaves.from, leaves.to FROM leaves WHERE employee_id = \$1 AND status = 'A' OR status='R' OR status = 'D' ORDER BY created_at DESC LIMIT 1;`
	mock.ExpectQuery(expectedQuery).WithArgs(check.EmployeeID).
		WillReturnRows(sqlmock.NewRows([]string{"from", "to"}).AddRow("10-10-2001", "20-10-2001"))

	applied, err := employeeDB.GetLastAppliedLeave(context.Background(), check)

	assert.NoError(t, err)

	assert.Equal(t, leave.From, applied.From)
	assert.Equal(t, leave.To, applied.To)

	err = mock.ExpectationsWereMet()

	if assert.NoError(t, err) {
		log.Println("Test1 Passed")
	}

	mock.ExpectQuery(expectedQuery).WithArgs(check.EmployeeID).WillReturnError(errors.New("no leaves found"))
	applied, err = employeeDB.GetLastAppliedLeave(context.Background(), check)

	assert.Error(t, err)

	if assert.EqualError(t, err, "no leaves found") {
		log.Println("Test 2 passed")
	}

}

func TestLeaveStatusHistory(t *testing.T) {

	mockDB, mock, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock SQL created successfully")
	}
	defer mockDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock SQL connected with GORM successfully")
	}

	// Create the employee database instance with the mock database
	employeeDB := NewEmployeeRepository(db)

	check := domain.Leave{
		EmployeeID: 1,
	}

	exceptedquery := `SELECT leave_type, leaves.to, leaves.from, status FROM leaves WHERE employee_id = \$1`
	rows := sqlmock.NewRows([]string{"leave_type", "to", "from", "status"}).
		AddRow("Casual", "20-10-2001", "10-10-2001", "A").
		AddRow("Sick", "2023-08-10", "2023-08-11", "P")

	mock.ExpectQuery(exceptedquery).WithArgs(check.EmployeeID).WillReturnRows(rows)

	history, err := employeeDB.LeaveStatusHistory(context.Background(), int(check.EmployeeID))

	assert.NoError(t, err)

	expectedHistory := []response.LeaveHistory{
		{
			Leave_type: "Casual",
			To:         "20-10-2001",
			From:       "10-10-2001",
			Status:     "A",
		},
		{
			Leave_type: "Sick",
			To:         "2023-08-10",
			From:       "2023-08-11",
			Status:     "P",
		},
	}
	assert.Equal(t, expectedHistory, history)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func TestAttendance(t *testing.T) {
	mockDb, mock, err := sqlmock.New()

	if assert.NoError(t, err) {
		log.Println("mi")
	}

	defer mockDb.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDb, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock SQL connected with GORM successfully")
	}

	employeeDB := NewEmployeeRepository(db)

	employeeID := 1

	expectedQuery := `select attendances.date, attendances.punch_in , attendances.punch_out, duties.duty_type from attendances inner join duties on attendances.employee_id = duties.employee_id where duties.status = 'C' and duties.employee_id = \$1;`
	rows := sqlmock.NewRows([]string{"date", "punch_in", "punch_out", "duty_type"}).
		AddRow("2023-06-26", "09:00:00", "17:00:00", "Regular").
		AddRow("2023-06-27", "09:30:00", "18:00:00", "Regular")

	mock.ExpectQuery(expectedQuery).WithArgs(employeeID).WillReturnRows(rows)

	attendance, err := employeeDB.Attendance(context.Background(), employeeID)
	if err != nil {
		t.Fatalf("Failed to get attendance: %v", err)
	}

	expectedAttendance := []response.Attendance{
		{Date: "2023-06-26", Punch_in: "09:00:00", Punch_out: "17:00:00", Duty_type: "Regular"},
		{Date: "2023-06-27", Punch_in: "09:30:00", Punch_out: "18:00:00", Duty_type: "Regular"},
	}

	if !reflect.DeepEqual(attendance, expectedAttendance) {
		t.Errorf("Mismatched attendance. Expected: %v, but got: %v", expectedAttendance, attendance)
	}

	// Ensure all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}

}
