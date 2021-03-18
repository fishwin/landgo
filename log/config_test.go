package log

import (
	"testing"
)

func init() {
	err := InitWithFile("config/log.yml")
	if err != nil {
		panic(err)
	}
}

func TestLogger(t *testing.T) {
	Logger().Sugar().Debug("1234")
	Logger().Sugar().Warn("1234")
	Logger().Sugar().Info("1234")
	Logger().Sugar().Error("1234")
	Logger().Sugar().Error("1234")
	Logger().Sugar().Error("1234")
	Logger().Sugar().Error("1234")

	Logger().Debug("1234")
	Logger().Warn("1234")
	Logger().Info("1234")
	Logger().Error("1234")
}
