package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/wybbb1/SiMo/internal/config"
	"github.com/wybbb1/SiMo/internal/log"
	user "github.com/wybbb1/SiMo/internal/model/user"
)

var DB *gorm.DB

func InitDB() {
	dbType := config.Config.Database.Type

	if dbType == "sqlite" {
		dbPath := config.Config.Database.Sqlite.Path
		
		dir := dbPath[:len(dbPath)-len("/ech0.db")] // 提取目录部分
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Logger.Fatal("Failed to create directory for database: " + err.Error())
		}

		var err error
		DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
		if err != nil {
			log.Logger.Fatal("Failed to connect to database: " + err.Error())
		}
	}

	if dbType == "mysql" {
		// Todo: MySQL support
	}

	if err := MigrateDB(); err != nil {
		log.Logger.Fatal("Failed to migrate database: " + err.Error())
	}
}

func CreateMysql() {
	// todo: CreateDB for MySQL
}

func MigrateDB() error {
	models := []interface{}{
		&user.User{},
	}

	return DB.AutoMigrate(models...)
}
