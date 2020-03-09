package log

import "go.uber.org/zap"

var logger *zap.Logger

func Logger() *zap.Logger {
	if logger == nil {
		panic("not init logger")
	}
	return logger
}
