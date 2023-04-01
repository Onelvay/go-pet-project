package model

type Config struct {
	HOST    string `yaml:"host"`
	USER    string `yaml:"user"`
	DB_NAME string `yaml:"dbname"`
	PORT    string `yaml:"port"`
	PASS    string `yaml:"pass"`
}

func NewConfig(h string, p string, d string, u string, ps string) *Config {
	return &Config{
		HOST:    h,
		PORT:    p,
		DB_NAME: d,
		USER:    u,
		PASS:    ps,
	}
}
