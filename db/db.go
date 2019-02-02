package db

import  "database/sql"


var DB *sql.DB

func SetDB(db *sql.DB) {
	DB = db
}


func GetDB() *sql.DB{
	return DB
}

func ConectDB() {
	var err error
	DB, err = sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"pharmacy_sh")
	if err != nil {
		panic(err.Error()) 
	}
}

func CloseDB() {
	defer DB.Close()
}
