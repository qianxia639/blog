package config

type Captcha struct {
	Height          int    `mapstructure:"height"`
	Width           int    `mapstructure:"width"`
	NoiseCount      int    `mapstructure:"noise_count"`
	ShowLineOptions int    `mapstructure:"show_line_options"`
	Length          int    `mapstructure:"length"`
	Source          string `mapstructure:"source"`
	Color           struct {
		R uint8 `mapstructure:"R"`
		G uint8 `mapstructure:"G"`
		B uint8 `mapstructure:"B"`
		A uint8 `mapstructure:"A"`
	} `mapstructure:"color"`
}
