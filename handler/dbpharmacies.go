package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func GetPharmacies(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var pharmacies []*model.Pharmacy
	var page model.Page
	var count int

	rows, err := db.DB.Table("pharmacy").Select("id, cif, address, phone_number, schedule, name, guard, mail, (SELECT COUNT(*) from pharmacy_sh.pharmacy) as count").Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		for rows.Next() {
			var id, numberPhone int
			var name, address, scheduler, cif, mail, guard string
			rows.Scan(&id, &cif, &address, &numberPhone, &scheduler, &name, &guard, &mail, &count)

			pharmacy := model.Pharmacy{
				ID:          id,
				Name:        name,
				PhoneNumber: numberPhone,
				Guard:       guard,
				Address:     address,
				Schedule:    scheduler,
				Cif:         cif,
				Mail:        mail,
			}
			pharmacies = append(pharmacies, &pharmacy)
			page = util.GetPage(count, intPage)
		}
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_PHARMACY)
	}

	pharmaciesResponse := model.PharmaciesResponse{
		Pharmacies: pharmacies,
		Page:       page,
		Response:   response,
	}

	output, err := json.Marshal(pharmaciesResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func GetPharmacy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	var response model.RequestResponse
	var pharmacy model.Pharmacy

	rows, err := db.DB.Table("pharmacy").Select("id, cif, address, phone_number, schedule, name, guard, mail").Where("id = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, numberPhone int
			var name, address, scheduler, cif, mail, guard string
			rows.Scan(&id, &cif, &address, &numberPhone, &scheduler, &name, &guard, &mail)

			pharmacyAux := model.Pharmacy{
				ID:          id,
				Cif:         cif,
				Address:     address,
				PhoneNumber: numberPhone,
				Schedule:    scheduler,
				Name:        name,
				Guard:       guard,
				Mail:        mail,
			}

			pharmacy = pharmacyAux
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_PHARMACY)
	}

	pharmacyResponse := model.PharmacyResponse{
		Pharmacy: &pharmacy,
		Response: response,
	}

	output, err := json.Marshal(pharmacyResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func CreatePharmacy(w http.ResponseWriter, r *http.Request) {
	var pharmacy model.Pharmacy
	var response model.RequestResponse
	var passwordHash string

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	passwordHash, err = util.HashPassword(pharmacy.Password)

	pharmacy.Password  = passwordHash

	db := db.DB.Table("pharmacy").Create(&pharmacy)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, util.TITLE_PHARMACY)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func UpdatePharmacy(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

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

	util.CheckErr(err)

	db := db.DB.Table("pharmacy").Where("id = ?", nID).Updates(&pharmacy)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, util.TITLE_PHARMACY)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func DeletePharmacy(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("pharmacy").Where("id = ?", nID).Delete(&model.Pharmacy{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleNotExistError(int(db.RowsAffected), error.DELETE, util.TITLE_PHARMACY)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}
