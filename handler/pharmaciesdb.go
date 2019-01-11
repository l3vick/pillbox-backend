package handler

import (
	"net/http"
	"fmt" 
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/gorilla/mux"
)



func GetPharmacies(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	page := vars["page"]

	var pharmacies []*model.Pharmacy
	selDB, err := dbConnector.Query("SELECT id, cif, street, number_phone, schedule, `name`, guard, account FROM med LIMIT" + page + ",10")
	if err != nil {
		panic(err.Error())
	}
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

	for selDB.Next() {
		var id, numberPhone, guard int
		var name, street, scheduler, cif, account string
		err = selDB.Scan(&id, &name, &numberPhone, &guard, &street, &scheduler, &cif, &account)
		if err != nil {
			panic(err.Error())
		}
		pharmacy := model.Pharmacy{
			ID:   			id,
			Name: 			name,
			NumberPhone:  	numberPhone,
			Guard:			guard,
			Street:			street,
			Schedule:		scheduler,
			Cif:			cif,
			Account:		account,
		}
		pharmacies = append(pharmacies, &pharmacy)
	}
	output, err := json.Marshal(pharmacies)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	w.Write(output)
}

func GetPharmacy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	selDB, err := dbConnector.Query("SELECT id, cif, street, number_phone, schedule, `name`, guard, account FROM pharmacy WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	pharmacy := model.Pharmacy{}
	for selDB.Next() {
		var id, numberPhone, guard int
		var name, street, scheduler, cif, account string
		err = selDB.Scan(&id, &name, &numberPhone, &guard, &street, &scheduler, &cif, &account)
		if err != nil {
			panic(err.Error())
		}
		pharmacy.ID = id
		pharmacy.Name = name
		pharmacy.NumberPhone = numberPhone
		pharmacy.Guard = guard
		pharmacy.Street = street
		pharmacy.Schedule = scheduler
		pharmacy.Cif = cif
		pharmacy.Account = account
	}

	output, err := json.Marshal(pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	w.Write(output)
}

func CreatePharmacy(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var pharmacy model.Pharmacy
	err = json.Unmarshal(b, &pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	w.Write(output)
	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`pharmacy` (`name`, `cif`, `street`, `number_phone`, `schedule`, `guard`, `password`, `account`)  VALUES('%s', '%s', '%s', '%d', '%s', '%d', '%s', '%s')", pharmacy.Name, pharmacy.Cif, pharmacy.Street, pharmacy.NumberPhone, pharmacy.Schedule, pharmacy.Guard, pharmacy.Password, pharmacy.Account)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdatePharmacy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var pharmacy model.Pharmacy
	err = json.Unmarshal(b, &pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`pharmacy` SET `name` = '%s', `cif` = '%s, `street` = '%s', `schedule` = '%s', `password` = '%s', `phone_number` = '%d', `guard` = '%d', `account` = `%s` WHERE (`id` = '%s)", pharmacy.Name, pharmacy.Cif, pharmacy.Street, pharmacy.Schedule, pharmacy.Password, pharmacy.NumberPhone, pharmacy.Guard, pharmacy.Account,nID)

	fmt.Println(query)
	update, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeletePharmacy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `rds_pharmacy`.`pharmacy` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}