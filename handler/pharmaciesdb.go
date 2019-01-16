package handler

import (
	"github.com/l3vick/go-pharmacy/util"
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

	var pharmacies []*model.PharmacyNotNull
	selDB, err := dbConnector.Query("SELECT id, cif, address, phone_number, schedule, `name`, guard, mail FROM med LIMIT" + page + ",10")
	if err != nil {
		panic(err.Error())
	}
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")

	for selDB.Next() {
		var id int
		var numberPhone, guard util.JsonNullInt64
		var name, address, scheduler, cif, mail util.JsonNullString
		err = selDB.Scan(&id, &name, &numberPhone, &guard, &address, &scheduler, &cif, &mail)
		if err != nil {
			panic(err.Error())
		}
		pharmacy := model.PharmacyNotNull{
			ID:   			id,
			Name: 			name,
			NumberPhone:  	numberPhone,
			Guard:			guard,
			Address:		address,
			Schedule:		scheduler,
			Cif:			cif,
			Mail:			mail,
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

	selDB, err := dbConnector.Query("SELECT id, cif, address, phone_number, schedule, `name`, guard, mail FROM pharmacy WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	pharmacy := model.PharmacyNotNull{}
	for selDB.Next() {
		var id int
		var numberPhone, guard util.JsonNullInt64
		var name, address, scheduler, cif, mail util.JsonNullString
		err = selDB.Scan(&id, &name, &numberPhone, &guard, &address, &scheduler, &cif, &mail)
		if err != nil {
			panic(err.Error())
		}
		pharmacy.ID = id
		pharmacy.Name = name
		pharmacy.NumberPhone = numberPhone
		pharmacy.Guard = guard
		pharmacy.Address = address
		pharmacy.Schedule = scheduler
		pharmacy.Cif = cif
		pharmacy.Mail = mail
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
	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`pharmacy` (`name`, `cif`, `address`, `phone_number`, `schedule`, `guard`, `password`, `mail`)  VALUES('%s', '%s', '%s', '%d', '%s', '%d', '%s', '%s')", pharmacy.Name, pharmacy.Cif, pharmacy.Address, pharmacy.NumberPhone, pharmacy.Schedule, pharmacy.Guard, pharmacy.Password, pharmacy.Mail)

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

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`pharmacy` SET  `cif` = '%s', `address` = '%s', `phone_number` = '%d', `schedule` = '%s', `name` = '%s', `guard` = '%d', `password` = '%s', `mail` = `%s` WHERE (`id` = '%s)", pharmacy.Cif, pharmacy.Address, pharmacy.NumberPhone, pharmacy.Schedule, pharmacy.Name, pharmacy.Guard, pharmacy.Password, pharmacy.Mail,nID)

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