package main

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Med struct {
	id   int
	name string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Estoy funcionando!!!!-------------------" + message
	w.Write([]byte(message))
}

func getmed(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"rds_pharmacy")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	selDB, err := db.Query("SELECT * FROM med ")
	if err != nil {
		panic(err.Error())
	}

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

	w.Header().Set("Content-Type", "application/json")

	for selDB.Next() {
		var id int
		var name string
		err = selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}
		/*
			e := Med{
				id:   &id,
				name: &name
			}
			medJSON, err := json.Marshal(e)
			if err != nil {
				// handle error
			}*/
		//message = "id: " + string(med.id) + " name: " + med.name
		message = " name " + name

		w.Write([]byte(message))
	}

	defer db.Close()

}

func conectDB() {
	db, err := sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"rds_pharmacy")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func main() {

	//conectDB()

	http.HandleFunc("/getmed", getmed)

	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
