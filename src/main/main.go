package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Med struct {
	ID   int
	Name string
	Pvp  int
}

var db *sql.DB

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Estoy funcionando!!!!-------------------" + message
	w.Write([]byte(message))
}

func getmed(w http.ResponseWriter, r *http.Request) {

	var meds []*Med

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

		med := Med{
			ID:   id,
			Name: name,
			Pvp:  pvp,
		}

		meds = append(meds, &med)

	}

	medJSON, err := json.MarshalIndent(meds, "", " ")
	if err != nil {
		// handle error
	}

	//	fmt.Println(meds)

	w.Write([]byte(medJSON))

	defer db.Close()

}

func getMedById(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM med WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	med := Med{}
	for selDB.Next() {
		var id, pvp int
		var name string
		err = selDB.Scan(&id, &name, &pvp)
		if err != nil {
			panic(err.Error())
		}
		med.ID = id
		med.Name = name
		med.Pvp = pvp
	}

	medJSON, err := json.MarshalIndent(med, "", " ")
	if err != nil {
		// handle error
	}

	//	fmt.Println(meds)

	w.Write([]byte(medJSON))

	defer db.Close()
}

func conectDB() {

	var err error

	db, err = sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"rds_pharmacy")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

}

func main() {

	conectDB()

	http.HandleFunc("/getmed", getmed)

	http.HandleFunc("/getmedbyid", getMedById)

	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
