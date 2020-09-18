package app

import (
	"github.com/benjcal/go-starter/app/models"
	dbPkg "github.com/benjcal/go-starter/pkg/db"
	"gorm.io/gorm"
)

type App struct {
	Config models.Config
	db     *gorm.DB
}

func Init(c models.Config) (App, error) {
	db, err := dbPkg.GetDB(c)
	if err != nil {
		return App{}, err
	}

	return App{
		Config: c,
		db:     db,
	}, nil
}
