package config

type Config struct {
	Http  Http  `mapstructure:"http"`
	MySQL MySQL `mapstructure:"mysql"`
	Qiniu Qiniu `mapstructure:"qiniu"`
}
