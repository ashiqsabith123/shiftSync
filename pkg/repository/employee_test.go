package repository

import (
	"context"
	"fmt"
	"log"
	"shiftsync/pkg/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddEmployee(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock sql created succesfully")
	}
	defer mockDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock sql connected with gorm succesfully")
	}
	defer mockDB.Close()

	emp := domain.Employee{
		First_name: "Ashiq",
		Last_name:  "Sabith",
		Email:      "ashiqsabith328@gmail.com",
		User_name:  "ashiq328",
		Pass_word:  "Ashiq@123",
		Phone:      8606863748,
	}

	query := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&emp)

	})
	fmt.Println(query)
	mock.ExpectBegin()

	mock.ExpectExec(query).
		WithArgs(emp.First_name, emp.Last_name, emp.Email, emp.Phone, emp.User_name, emp.Pass_word).
		WillReturnResult(sqlmock.NewResult(1, 1))

	employeeDB := NewEmployeeRepository(db)

	err = employeeDB.AddEmployee(context.Background(), emp)
	assert.NoError(t, err)
}
