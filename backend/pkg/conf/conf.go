package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/benjcal/go-starter/app/models"
)

func GetConfig(confString string) (c models.Config, err error) {
	_, err = toml.Decode(confString, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
