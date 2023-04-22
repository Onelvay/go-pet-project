package postgres

import (
	"fmt"

	"github.com/Onelvay/docker-compose-project/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
func NewPostgresDb(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.HOST, cfg.PORT, cfg.USER, cfg.DB_NAME, cfg.PASS)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.Book{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{}, &domain.Refresh_token{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{}, &domain.Order{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.Order{}, &domain.FinalResponse{})

	if err != nil {
		panic(err)
	}

	return db
}
