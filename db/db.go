package db

import  (
	"github.com/jinzhu/gorm"
	
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConectDB() {
	var err error
	DB, err = gorm.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"pharmacy_sh?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error()) 
	}
}

func CloseDB() {
	defer DB.Close()
}
