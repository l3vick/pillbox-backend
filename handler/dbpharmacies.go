package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPharmacies(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage)

	query := fmt.Sprintf("SELECT id, cif, address, phone_number, schedule, `name`, guard, mail, (SELECT COUNT(*) from pharmacy_sh.pharmacy) as count FROM pharmacy_sh.pharmacy LIMIT " + elem + ",10")

	fmt.Println(query)
	var pharmacies []*model.Pharmacy

	var page model.Page

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, numberPhone, guard, count int
		var name, address, scheduler, cif, mail string
		err = selDB.Scan(&id, &cif, &address, &numberPhone, &scheduler, &name, &guard, &mail, &count)
		if err != nil {
			panic(err.Error())
		}
		pharmacy := model.Pharmacy{
			ID:          id,
			Name:        name,
			NumberPhone: numberPhone,
			Guard:       guard,
			Address:     address,
			Schedule:    scheduler,
			Cif:         cif,
			Mail:        mail,
		}
		pharmacies = append(pharmacies, &pharmacy)
		page = util.GetPage(count, intPage)
	}

	response := model.PharmacyResponse{
		Pharmacy: pharmacies,
		Page:     page,
	}

	output, err := json.Marshal(response)
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

	pharmacy := model.Pharmacy{}
	for selDB.Next() {
		var id, numberPhone, guard int
		var name, address, scheduler, cif, mail string
		err = selDB.Scan(&id, &cif, &address, &numberPhone, &scheduler, &name, &guard, &mail)
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
	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`pharmacy` (`name`, `cif`, `address`, `phone_number`, `schedule`, `guard`, `password`, `mail`)  VALUES('%s', '%s', '%s', '%d', '%s', '%d', '%s', '%s')", pharmacy.Name, pharmacy.Cif, pharmacy.Address, pharmacy.NumberPhone, pharmacy.Schedule, pharmacy.Guard, pharmacy.Password, pharmacy.Mail)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)

	var pharmacyResponse model.RequestResponse
	if err != nil {
		pharmacyResponse.Code = 500
		pharmacyResponse.Message = err.Error()
	} else {
		pharmacyResponse.Code = 200
		pharmacyResponse.Message = "Pharmacy creada con éxito"
	}

	output, err2 := json.Marshal(pharmacyResponse)
	if err2 != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

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

	var query string = fmt.Sprintf("UPDATE `pharmacy_sh`.`pharmacy` SET  `cif` = '%s', `address` = '%s', `phone_number` = '%d', `schedule` = '%s', `name` = '%s', `guard` = '%d', `password` = '%s', `mail` = '%s' WHERE (`id` = '%s')", pharmacy.Cif, pharmacy.Address, pharmacy.NumberPhone, pharmacy.Schedule, pharmacy.Name, pharmacy.Guard, pharmacy.Password, pharmacy.Mail, nID)

	fmt.Println(query)
	update, err := dbConnector.Query(query)

	var pharmacyResponse model.RequestResponse
	if err != nil {
		pharmacyResponse.Code = 500
		pharmacyResponse.Message = err.Error()
	} else {
		pharmacyResponse.Code = 200
		pharmacyResponse.Message = "Pharmacy actualizada con éxito"
	}

	output, err2 := json.Marshal(pharmacyResponse)
	if err2 != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer update.Close()
}

func DeletePharmacy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`pharmacy` WHERE (`id` = '%s')", nID)

	fmt.Println(query)

	insert, err := dbConnector.Query(query)

	var pharmacyResponse model.RequestResponse
	if err != nil {
		pharmacyResponse.Code = 500
		pharmacyResponse.Message = err.Error()
	} else {
		pharmacyResponse.Code = 200
		pharmacyResponse.Message = "Pharmacy borrada con éxito"
	}

	output, err := json.Marshal(pharmacyResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer insert.Close()
}
