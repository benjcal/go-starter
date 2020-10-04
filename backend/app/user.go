package app

import (
	"fmt"
	"github.com/benjcal/go-starter/pkg/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func (a App) UserRegister(email, pwd, name string) error {
	err := a.db.Exec("SELECT api.add_user(?,?,?)", email, pwd, name).Error
	if err != nil {
		return err
	}

	return nil
}

func (a App) UserLogin(email, pwd string) (map[string]interface{}, error) {
	v := map[string]interface{}{}

	err := a.db.Raw("SELECT * FROM api.valid_user(?,?)", email, pwd).Scan(&v).Error
	if err != nil && !strings.Contains(err.Error(), schema.ErrUnsupportedDataType.Error()) {
		return nil, err
	}

	_, ok := v["id"]
	if !ok {
		return nil, fmt.Errorf("no user found")
	}

	jwtAuth := auth.GetJWTAuth(a.Config)

	_, jwt, _ := jwtAuth.Encode(jwt.MapClaims{"user_id": 33, "exp": jwtauth.ExpireIn(time.Minute)})

	v["token"] = jwt

	return v, nil
}
