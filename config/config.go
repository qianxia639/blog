package config

type Config struct {
	Http    Http    `mapstructure:"http"`
	MySQL   MySQL   `mapstructure:"mysql"`
	Redis   Redis   `mapstructure:"redis"`
	Qiniu   Qiniu   `mapstructure:"qiniu"`
	Captcha Captcha `mapstructure:"captcha"`
}
