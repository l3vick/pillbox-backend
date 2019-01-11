package db

import  "database/sql"


var db *sql.DB


func GetDB() *sql.DB{
	return db
}

func ConectDB() {
	var err error
	db, err = sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"rds_pharmacy")
	if err != nil {
		panic(err.Error()) 
	}
}

func CloseDB() {
	defer db.Close()
}
