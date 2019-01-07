package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/src/main/model/med"
)

type Users struct {
	Id              int
	Name            string
	MedBreackfast   int
	MedLaunch       int
	MedDinner       int
	AlarmBreackfast string
	AlarmLaunch     string
	AlarmDinner     string
}

var db *sql.DB

func root(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "App Farmacias" + message
	w.Write([]byte(message))
}

func conectDB() {
	var err error
	db, err = sql.Open("mysql", "rds_pharmacy_00"+":"+"phar00macy"+"@tcp("+"rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com:3306"+")/"+"rds_pharmacy")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
}

func GetMeds(w http.ResponseWriter, r *http.Request) {
	conectDB()
	var meds []*med.Med
	selDB, err := db.Query("SELECT * FROM med LIMIT 10")
	if err != nil {
		panic(err.Error())
	}
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

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
	output, err := json.Marshal(meds)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	defer db.Close()
}

func GetMed(w http.ResponseWriter, r *http.Request) {
	conectDB()

	vars := mux.Vars(r)

	nID := vars["id"]

	selDB, err := db.Query("SELECT * FROM med WHERE id=?", nID)
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

	//medJSON, err := json.MarshalIndent(med, "", " ")
	//if err != nil {
	// handle error
	//}

	//	w.Write([]byte(medJSON))

	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	defer db.Close()
}

func CreateMed(w http.ResponseWriter, r *http.Request) {
	conectDB()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var t Med
	err = json.Unmarshal(b, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(t.Name)
	output, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	var query string = fmt.Sprintf("INSERT INTO `rds_pharmacy`.`med` (`name`, `pvp`) VALUES('%s','%d')", t.Name, t.Pvp)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	//w.Write([]byte(medJSON))
	defer insert.Close()
	defer db.Close()

}

func UpdateMed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	nID := vars["id"]

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var t Med
	err = json.Unmarshal(b, &t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(t.Name)
	output, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`med` SET `name` = '%s', `pvp` = '%d' WHERE (`id` = '%s')", t.Name, t.Pvp, nID)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	defer db.Close()
}

func DeleteMed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nID := vars["id"]

	var query string = fmt.Sprintf("DELETE FROM `rds_pharmacy`.`med` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
	defer db.Close()
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
			Id:              id,
			Name:            name,
			MedBreackfast:   medBreackfast,
			MedLaunch:       medLaunch,
			MedDinner:       medDinner,
			AlarmBreackfast: alarmBreackfast,
			AlarmLaunch:     alarmLaunch,
			AlarmDinner:     alarmDinner,
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

//GetUser ...
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
	router.HandleFunc("/", root).Methods("GET")
	router.HandleFunc("/meds", GetMeds).Methods("GET")
	router.HandleFunc("/meds/{id}", GetMed).Methods("GET")
	router.HandleFunc("/meds", CreateMed).Methods("POST")
	router.HandleFunc("/meds/{id}", UpdateMed).Methods("PUT")
	router.HandleFunc("/meds/{id}", DeleteMed).Methods("DELETE")

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
