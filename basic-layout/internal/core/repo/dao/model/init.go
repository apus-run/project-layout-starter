package model

import (
	"gorm.io/gorm"
)

func InitTables(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &User{})
	if err != nil {
		panic(err)
	}
}
