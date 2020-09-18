package models

type Config struct {
	DBDSN     string `toml:"db_dsn"`
	JWTSecret string `toml:"jwt_secret"`
}
