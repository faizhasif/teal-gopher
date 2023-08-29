package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	db_url, db_port, db_name, db_user, db_pass := os.Getenv("DB_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASS")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass, db_url, db_port, db_name)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}
}
