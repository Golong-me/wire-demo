package conf

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ZapProvider = wire.NewSet(NewZap)

func NewZap() *zap.Logger {
	logger := zap.NewExample()
	return logger
}
