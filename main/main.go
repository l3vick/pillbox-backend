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
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	MedBreakfast   string `json:"med_breakfast"`
	MedLaunch      string `json:"med_launch"`
	MedDinner      string `json:"med_dinner"`
	AlarmBreakfast string `json:"alarm_breakfast"`
	AlarmLaunch    string `json:"alarm_launch"`
	AlarmDinner    string `json:"alarm_dinner"`
	Password       string `json:"password"`
	IDPharmacy     int    `json:"id_pharmacy"`
}

type Med struct {
	ID   int
	Name string `json:"name"`
	Pvp  int    `json:"pvp"`
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

func closeDB() {
	defer db.Close()
}

func GetMeds(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var meds []*Med
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
}

func GetMed(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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

	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func CreateMed(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var med Med
	err = json.Unmarshal(b, &med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(med.Name)
	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`med` (`name`, `pvp`) VALUES('%s','%d')", med.Name, med.Pvp)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateMed(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var med Med
	err = json.Unmarshal(b, &med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`med` SET `name` = '%s', `pvp` = '%d' WHERE (`id` = '%s')", med.Name, med.Pvp, nID)

	fmt.Println(query)
	update, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeleteMed(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)

	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `rds_pharmacy`.`med` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var users []*User

	selDB, err := db.Query("SELECT * FROM users LIMIT 10")
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	for selDB.Next() {
		var id, idPharmacy int
		var name, medBreakfast, medLaunch, medDinner, alarmBreakfast, alarmLaunch, alarmDinner, password string
		err = selDB.Scan(&id, &name, &medBreakfast, &medLaunch, &medDinner, &alarmBreakfast, &alarmLaunch, &alarmDinner, &password, &idPharmacy)

		if err != nil {
			panic(err.Error())
		}

		user := User{
			ID:             id,
			Name:           name,
			MedBreakfast:   medBreakfast,
			MedLaunch:      medLaunch,
			MedDinner:      medDinner,
			AlarmBreakfast: alarmBreakfast,
			AlarmLaunch:    alarmLaunch,
			AlarmDinner:    alarmDinner,
			Password:       password,
			IDPharmacy:     idPharmacy,
		}

		users = append(users, &user)

	}

	usersJSON, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		panic(err.Error())
	}

	w.Write([]byte(usersJSON))
}

//GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	nID := r.URL.Query().Get("id")
	user := User{}

	selDB, err := db.Query("SELECT * FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy int
		var name, medBreakfast, medLaunch, medDinner, alarmBreakfast, alarmLaunch, alarmDinner, password string
		err = selDB.Scan(&id, &name, &medBreakfast, &medDinner, &medLaunch, &alarmDinner, &alarmLaunch, &alarmBreakfast, &password, &idPharmacy)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.MedBreakfast = medBreakfast
		user.MedLaunch = medLaunch
		user.MedDinner = medDinner
		user.AlarmBreakfast = alarmBreakfast
		user.AlarmLaunch = alarmLaunch
		user.AlarmDinner = alarmDinner
		user.IDPharmacy = idPharmacy
		user.Password = password
	}

	userJSON, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		// handle error
	}

	w.Write([]byte(userJSON))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`users` (`name`, `med_breakfast`, `med_launch`, `med_dinner`, `alarm_breakfast`, `alarm_launch`, `alarm_dinner`, `password`, `id_pharmacy`)  VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d')", user.Name, user.MedBreakfast, user.MedLaunch, user.MedDinner, user.AlarmBreakfast, user.AlarmLaunch, user.AlarmDinner, user.Password, user.IDPharmacy)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`users` SET `name` = '%s', `med_breakfast` = '%s', `med_launch` = '%s', `med_dinner` = '%s', `alarm_breakfast` = '%s', `alarm_launch` = '%s', `alarm_dinner` = '%s', `id_pharmacy` = '%d' WHERE (`id` = '%s')", user.Name, user.MedBreakfast, user.MedLaunch, user.MedDinner, user.AlarmBreakfast, user.AlarmBreakfast, user.AlarmBreakfast, user.IDPharmacy, nID)

	fmt.Println(query)
	update, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)

	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `rds_pharmacy`.`users` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
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
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/pharmacies", GetPharmacies).Methods("GET")
	router.HandleFunc("/pharmacies/{id}", GetPharmacy).Methods("GET")
	router.HandleFunc("/pharmacies", CreatePharmacy).Methods("POST")
	router.HandleFunc("/pharmacies", UpdatePharmacy).Methods("PUT")
	router.HandleFunc("/pharmacies/{id}", DeletePharmacy).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}
