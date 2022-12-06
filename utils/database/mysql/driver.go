package mysql

import (
	"fmt"
	"log"

	"github.com/GP2-Group5/Backend/config"
	userdata "github.com/GP2-Group5/Backend/feature/users/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(userdata.Users{})
}
