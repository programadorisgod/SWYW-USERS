package logger

import (
	"swyw-users/src/utils/enviroment"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {

	appMode := enviroment.GetAppMode()

	if appMode == enviroment.ProdMode {
		Log, _ = zap.NewProduction()
		return
	}

	Log, _ = zap.NewDevelopment()

}
