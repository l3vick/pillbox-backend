package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func CreateMed(w http.ResponseWriter, r *http.Request) {
	var med model.Med
	var response model.RequestResponse

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("med").Create(&med)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, util.TITLE_MED)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func GetMed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	var response model.RequestResponse
	var med model.Med

	rows, err := db.DB.Table("med").Select("*").Where("id = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, pharmacyId int
			var name, description string
			rows.Scan(&id, &name, &description, &pharmacyId)

			medAux := model.Med{
				ID:          id,
				Name:        name,
				Description: description,
				IDPharmacy:  pharmacyId,
			}
			med = medAux
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_MED)
	}

	medResponse := model.MedResponse{
		Med:      &med,
		Response: response,
	}

	output, err := json.Marshal(medResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func UpdateMed(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var med model.Med

	err = json.Unmarshal(b, &med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("med").Where("id = ?", nID).Updates(&med)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, util.TITLE_MED)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func DeleteMed(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("med").Where("id = ?", nID).Delete(&model.Med{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleNotExistError(int(db.RowsAffected), error.DELETE, util.TITLE_MED)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}

func GetMeds(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name != "" {
		FilterMedByName(name, w, r)
	} else {
		GetAllMeds(w, r)
	}
}

func GetAllMeds(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var meds []*model.Med
	var page model.Page
	var count int

	rows, err := db.DB.Table("med").Select("id, name, description, id_pharmacy,(SELECT COUNT(*)  from pharmacy_sh.med) as count").Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		for rows.Next() {
			var id, pharmacyID int
			var name, description string
			rows.Scan(&id, &name, &description, &pharmacyID, &count)

			med := model.Med{
				ID:          id,
				Name:        name,
				Description: description,
				IDPharmacy:  pharmacyID,
			}
			meds = append(meds, &med)

		}
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_MED)
		page = util.GetPage(count, intPage)
	}

	medsResponse := model.MedsResponse{
		Meds:     meds,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func FilterMedByName(name string, w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var meds []*model.Med
	var page model.Page

	rows, err := db.DB.Table("med").Select("*").Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, pharmacyID int
			var name, description string
			rows.Scan(&id, &name, &description, &pharmacyID)

			medAux := model.Med{
				ID:          id,
				Name:        name,
				Description: description,
				IDPharmacy:  pharmacyID,
			}
			meds = append(meds, &medAux)
		}
		page = util.GetPage(count, intPage)
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_MED)
	}

	medsResponse := model.MedsResponse{
		Meds:     meds,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func GetMedsByPharmacyID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var meds []*model.Med
	var page model.Page
	var count int

	rows, err := db.DB.Table("med").Select("id, name, description, id_pharmacy,(SELECT COUNT(*)  from pharmacy_sh.med WHERE id_pharmacy = "+nID+") as count").Where("id_pharmacy = ?", nID).Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		for rows.Next() {
			var id, pharmacyID int
			var name, description string
			rows.Scan(&id, &name, &description, &pharmacyID, &count)

			med := model.Med{
				ID:          id,
				Name:        name,
				Description: description,
				IDPharmacy:  pharmacyID,
			}
			meds = append(meds, &med)
			page = util.GetPage(count, intPage)
		}
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_MED)
	}

	medsResponse := model.MedsResponse{
		Meds:     meds,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}
