package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Med struct {
	ID   int
	Name string
	Pvp  int
}

type Users struct {
	Id               int
	Name             string
	MedBreackfast   int
	MedLaunch       int
	MedDinner       int
	AlarmBreackfast string
	AlarmLaunch     string
	AlarmDinner     string
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []*Users

	selDB, err := db.Query("SELECT * FROM users LIMIT 10")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	for selDB.Next() {
		var id int
		var name string
		var medBreackfast int
		var medLaunch int
		var medDinner int
		var alarmBreackfast string
		var alarmLaunch string
		var alarmDinner string

		err = selDB.Scan(&id, &name, &medBreackfast, &medLaunch, &medDinner, &alarmBreackfast, &alarmLaunch, &alarmDinner)
		
		if err != nil {
			panic(err.Error())
		}

		user := Users{
			Id: id,
			Name: name,
			MedBreackfast: medBreackfast,
			MedLaunch: medLaunch,
			MedDinner: medDinner,
			AlarmBreackfast: alarmBreackfast,
			AlarmLaunch: alarmLaunch,
			AlarmDinner: alarmDinner,
		}

		users = append(users, &user)

	}

	usersJSON, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		// handle error
	}

	w.Write([]byte(usersJSON))

	defer db.Close()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	user := Users{}

	selDB, err := db.Query("SELECT * FROM users WHERE id=?", nId)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, medBreackfast, medLaunch, medDinner int
		var name, alarmBreackfast, alarmLaunch, alarmDinner string
		err = selDB.Scan(&id, &name, &medBreackfast, &medDinner, &medLaunch, &alarmDinner, &alarmLaunch, &alarmBreackfast)

		if err != nil {
			panic(err.Error())
		}

		user.Id = id
		user.Name = name
		user.MedBreackfast = medBreackfast
		user.MedLaunch = medLaunch
		user.MedDinner = medDinner
		user.AlarmBreackfast = alarmBreackfast
		user.AlarmLaunch = alarmLaunch
		user.AlarmDinner = alarmDinner
	}

	userJSON, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		// handle error
	}

	w.Write([]byte(userJSON))

	defer db.Close()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func main() {

	conectDB()

	router := mux.NewRouter()
	router.HandleFunc("/", sayHello).Methods("GET")

	http.HandleFunc("/getmed", getmed)

	http.HandleFunc("/getmedbyid", getMedById)

	//router.HandleFunc("/meds", getmed).Methods("GET")
	//router.HandleFunc("/meds/{id}", getMedById).Methods("GET")
	//router.HandleFunc("/meds/{id}", CreateMed).Methods("POST")
	//router.HandleFunc("/meds/{id}", DeleteMed).Methods("DELETE")

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
