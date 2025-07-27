package database

import (
	"fmt"
	"log"

	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.GetEnv("DB_USER", ""), env.GetEnv("DB_PASSWORD", ""), env.GetEnv("DB_HOST", "127.0.0.1"), env.GetEnv("DB_PORT", "3306"), env.GetEnv("DB_NAME", ""))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database \n", err.Error())
	}

	err = DB.AutoMigrate(&models.User{}, &models.UserSession{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)
}
