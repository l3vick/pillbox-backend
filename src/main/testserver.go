package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Med struct {
	ID   int
	Name string
	Pvp  int
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
	selDB, err := db.Query("SELECT * FROM med LIMIT 10")
	if err != nil {
		panic(err.Error())
	}

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

	w.Header().Set("Content-Type", "application/json")

	for selDB.Next() {
		var id int
		var name string
		var pvp int
		err = selDB.Scan(&id, &name, &pvp)
		if err != nil {
			panic(err.Error())
		}

		e := Med{
			ID:   id,
			Name: name,
			Pvp:  pvp,
		}
		medJSON, err := json.Marshal(e)
		if err != nil {
			// handle error
		}

		fmt.Println(e)

		w.Write([]byte(medJSON))
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
