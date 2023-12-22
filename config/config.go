package config

import (
	"sing-box-cli/internal/utils"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BaseDir string `env:"BASE_DIR" env-default:"/data/sing-box"`
}

var Cfg Config

func init() {
	utils.Check(cleanenv.ReadEnv(&Cfg))
}
