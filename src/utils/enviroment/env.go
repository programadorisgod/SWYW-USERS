package enviroment

import "os"

type AppMode string

const (
	DevMode  AppMode = "dev"
	ProdMode AppMode = "prod"
)

func GetAppMode() AppMode {
	env := os.Getenv("APP_ENV")

	if env == "prod" {
		return ProdMode
	}

	return DevMode
}
