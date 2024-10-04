package conf

import "go.uber.org/zap"

var Logger *zap.Logger

func NewZap() {
	logger := zap.NewExample()
	Logger = logger
}
