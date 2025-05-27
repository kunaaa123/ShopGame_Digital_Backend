package mysql

import (
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/kunaaa123/SHOPGAME/internal/configs"
	"github.com/kunaaa123/SHOPGAME/internal/models"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var instance *gorm.DB

func Init(configs *configs.Config) error {
	mysqlConfig := mysql.Config{
		User:                 configs.Server.MySQLUsername,
		Passwd:               configs.Server.MySQLPassword,
		Net:                  "tcp",
		Addr:                 configs.Server.MySQLHost,
		DBName:               configs.Server.MySQLDatabase,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		ParseTime:            true,
		MultiStatements:      true,
	}

	dsn := mysqlConfig.FormatDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	gormConfig := &gorm.Config{
		Logger: newLogger,
	}

	db, err := gorm.Open(driver.Open(dsn), gormConfig)
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migration for local test
	err = db.AutoMigrate(&models.GameDB{})
	if err != nil {
		panic("failed to auto migrate")
	}

	instance = db

	log.Printf("connected to database successfully!")

	return nil
}

func GetDB(configs *configs.Config) *gorm.DB {
	if instance != nil {
		return instance
	}

	if err := Init(configs); err != nil {
		log.Panicf("cannot init mysql: %v", err)
	}

	db, err := instance.DB()
	if err != nil {
		log.Fatalf("failed to get DB from GORM: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	return instance
}
