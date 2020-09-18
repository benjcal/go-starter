package db

import (
	"github.com/benjcal/go-starter/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB(c models.Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(c.DBDSN), nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
