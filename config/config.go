package config

import (
	"context"

	"github.com/kkyr/fig"
)

type ctxKey struct{}

type Telegram struct {
	Token string `fig:"token" yaml:"token"`
	User  int64  `fig:"user" yaml:"user"`
}

type Paths struct {
	BkpDir string `fig:"bkp_dir" yaml:"bkp_dir"`
}

type AppConfig struct {
	Telegram Telegram `fig:"telegram" yaml:"telegram"`
	Paths    Paths    `fig:"paths" yaml:"paths"`
}

var current AppConfig

func Ctx(ctx context.Context) AppConfig {
	cf, _ := ctx.Value(ctxKey{}).(AppConfig)

	fig.Load(&cf,
		fig.File("config.yml"),
		fig.Dirs("."),
	)
	return cf
}

func SetConfig(cfg AppConfig) {
	current = cfg
}
