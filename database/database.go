package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

func initialMigration(db *gorm.DB) {
	db.AutoMigrate(&Sample{})
}

func NewDatabase() (*gorm.DB, error) {
	conn, err := gorm.Open("sqlite3", os.Getenv("DB_FILE"))
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	initialMigration(conn)
	return conn, nil
}
