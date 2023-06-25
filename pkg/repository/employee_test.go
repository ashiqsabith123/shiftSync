package repository

import (
	"fmt"
	"log"
	"reflect"
	"shiftsync/pkg/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddEmployee(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if assert.NoError(t, err) {
		log.Println("Mock sql created succesfully")

	}
	defer mockDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB, DriverName: "postgres"}), &gorm.Config{})
	if assert.NoError(t, err) {
		log.Println("Mock sql connected with gorm succesfully")
	}

	// config, _ := config.LoadConfig()
	// db.ConnectToDatbase(config)

	// DB := db.GetDatabaseInstance()

	emp := domain.Employee{
		First_name: "Ashiq",
		Last_name:  "Sabith",
		Email:      "ashiqsabith328@gmail.com",
		User_name:  "ashiq328",
		Pass_word:  "Ashiq@123",
		Phone:      8606863748,
	}

	// query := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Create(&emp)

	// })

	stmt := db.Create(&emp).Statement
	// v := reflect.ValueOf(stmt)
	// tv := v.Type()

	// for i := 0; i < v.NumField(); i++ {
	// 	field := v.Field(i)
	// 	fieldName := tv.Field(i).Name
	// 	fieldValue := field.Interface()

	// 	fmt.Printf("%s: %v\n", fieldName, fieldValue)
	// }

	qu := db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	fmt.Println(qu)

	fmt.Println("q", query)
	mock.ExpectBegin()
	mock.ExpectRollback()
	mock.ExpectExec("").WillReturnResult(sqlmock.NewResult(1, 1))

	// // // mock.ExpectCommit()
	// // mock.ExpectRollback()

	// employeeDB := NewEmployeeRepository(db)

	// eerr := employeeDB.AddEmployee(context.Background(), emp)
	// assert.NoError(t, eerr)

	//assert.NoError(t, mock.ExpectationsWereMet())
}
