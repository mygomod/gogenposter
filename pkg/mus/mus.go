package mus

import (
	"github.com/mygomod/muses/pkg/logger"
	"github.com/mygomod/muses/pkg/oss"
)

var (
	Logger *logger.Client
	Oss    *oss.Client
)

func Init() error {
	Logger = logger.Caller("system")
	Oss = oss.Caller("poster")
	return nil
}
