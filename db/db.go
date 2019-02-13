package db

import (
	"github.com/jinzhu/gorm"
	"github.com/l3vick/go-pharmacy/keys"
)

var DB *gorm.DB

func ConectDB() {
	var err error
	DB, err = gorm.Open(keys.DB_TYPE, keys.DB_USER+":"+keys.DB_PASSWORD+"@tcp("+keys.DB_DNS+")/"+keys.DB_NAME_SPECS)
	if err != nil {
		panic(err.Error())
	}
}

func CloseDB() {
	defer DB.Close()
}
