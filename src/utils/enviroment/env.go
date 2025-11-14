package enviroment

import (
	"swyw-users/src/config"
)

type AppMode string

const (
	DevMode  AppMode = "dev"
	ProdMode AppMode = "prod"
)

func GetAppMode() AppMode {
	env := config.Env.AppEnv

	if env == "prod" {
		return ProdMode
	}

	return DevMode
}
