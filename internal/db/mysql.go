package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/kunaaa123/SHOPGAME/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var (
	instance *gorm.DB
	once     sync.Once
)

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		DBName:   "gameshop",
	}
}

func GetDB(config *DBConfig) *gorm.DB {
	once.Do(func() {
		if os.Getenv("GO_ENV") == "test" {
			config.DBName = "gameshop_test"
		}

		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DBName,
		)

		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			log.Fatal("เชื่อมต่อฐานข้อมูลไม่ผ่าน", err)
		}
		err = db.AutoMigrate(&domain.Game{})
		if err != nil {
			log.Fatal("ไม่สามารถสร้างตารางได้", err)
		}
		instance = db
	})
	return instance
}
