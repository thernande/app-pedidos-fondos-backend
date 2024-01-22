package database

import (
	"fmt"
	"log"

	"github.com/thernande/app-pedidos-fondos-backend/internal/config/enviroment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Host     string
	Port     uint
	User     string
	Password string
	DbName   string
}

func New() *Database {
	env := enviroment.New()
	env.Load()
	return &Database{
		Host:     env.DBHost,
		Port:     env.DBPort,
		User:     env.DBUser,
		Password: env.DBPass,
		DbName:   env.DBName,
	}
}

func (d *Database) MySQL() *gorm.DB {
	env := enviroment.New()
	env.Load()
	logType := logger.Default.LogMode(logger.Silent)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Password, d.Host, d.Port, d.DbName)
	if env.DBMode == "debug" {
		logType = logger.Default.LogMode(logger.Info)
	}
	if env.DBMode == "release" {
		logType = logger.Default.LogMode(logger.Silent)
	}
	Db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		PrepareStmt: true,
		QueryFields: true,
		Logger:      logType,
	})
	if err != nil {
		log.Println(err, "error connecting to database")
		return nil
	}

	return Db
}
