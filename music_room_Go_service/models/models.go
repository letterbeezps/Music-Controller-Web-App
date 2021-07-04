package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "./data/test.db")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Room{}, &SpotifyToken{})

}

func CloseDB() {
	defer db.Close()
}
