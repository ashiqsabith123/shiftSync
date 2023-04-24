package db

import (
	"fmt"
	"log"
	"shiftsync/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatbase(config config.Config) (*gorm.DB, error) {
	connstr := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", config.Db.DbHost, config.Db.DbUser, config.Db.DbName, config.Db.DbPort, config.Db.DbPaswword)
	db, err := gorm.Open(postgres.Open(connstr))

	if err != nil {
		log.Fatal("Failed to connect database")
		return nil, err
	}

	return db, nil
}
