package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func GetMeds(w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage)

	query := fmt.Sprintf("SELECT id, name, description, id_pharmacy,(SELECT COUNT(*)  from pharmacy_sh.med) as count FROM pharmacy_sh.med LIMIT " + elem + ",10")

	fmt.Println(query)

	var meds []*model.Med

	var page model.Page

	selDB, err := db.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, pharmacyID, count int
		var name, description string
		err = selDB.Scan(&id, &name, &description, &pharmacyID, &count)
		if err != nil {
			panic(err.Error())
		}

		med := model.Med{
			ID:          id,
			Name:        name,
			Description: description,
			PharmacyID:  pharmacyID,
		}
		meds = append(meds, &med)
		page = util.GetPage(count, intPage)
	}

	response := model.MedResponse{
		Meds: meds,
		Page: page,
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(output)
}

func GetMed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	nID := vars["id"]

	query := fmt.Sprintf("SELECT * FROM pharmacy_sh.med WHERE id = " + nID + "")

	fmt.Println(query)

	selDB, err := db.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	med := model.Med{}
	for selDB.Next() {
		var id, pharmacyId int
		var name, description string
		err = selDB.Scan(&id, &name, &description, &pharmacyId)
		if err != nil {
			panic(err.Error())
		}
		med.ID = id
		med.Name = name
		med.Description = description
		med.PharmacyID = pharmacyId
	}

	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(output)
}

func CreateMed(w http.ResponseWriter, r *http.Request) {

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

	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`med` (`name`, `description`, `id_pharmacy`) VALUES ('%s', '%s','%d' )", med.Name, med.Description, med.PharmacyID)

	fmt.Println(query)
	insert, err := db.DB.Query(query)

	var medsResponse model.RequestResponse
	if err != nil {
		medsResponse.Code = 500
		medsResponse.Message = err.Error()
	} else {
		medsResponse.Code = 200
		medsResponse.Message = "Med creado con éxito"
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer insert.Close()
}

func UpdateMed(w http.ResponseWriter, r *http.Request) {

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

	var query string = "UPDATE `pharmacy_sh`.`med` SET"

	if med.Name != "" {
		query = query + fmt.Sprintf("`name` = '%s'", med.Name)
	}

	if med.Name != "" && med.Description != "" {
		query = query + " , "
	}

	if med.Description != "" {
		query = query + fmt.Sprintf("`description` = '%s'", med.Description)
	}

	query = query + fmt.Sprintf(" WHERE (`id` = '%s')", nID)

	fmt.Println(query)

	update, err := db.DB.Query(query)

	var medsResponse model.RequestResponse
	if err != nil {
		medsResponse.Code = 500
		medsResponse.Message = err.Error()
	} else {
		medsResponse.Code = 200
		medsResponse.Message = "Med actualizado con éxito"
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer update.Close()
}

func DeleteMed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`med` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := db.DB.Query(query)

	var medsResponse model.RequestResponse
	if err != nil {
		medsResponse.Code = 500
		medsResponse.Message = err.Error()
	} else {
		medsResponse.Code = 200
		medsResponse.Message = "Med borrado con éxito"
	}

	output, err := json.Marshal(medsResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer insert.Close()
}
