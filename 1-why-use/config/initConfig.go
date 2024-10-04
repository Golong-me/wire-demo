package config

import "wire-why/internal/conf"

func InitAllConfig() {
	conf.NewZap()
	conf.NewGorm()
}
