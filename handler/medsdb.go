package handler

import (
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/nullsql"
	"github.com/l3vick/go-pharmacy/util"
	"github.com/gorilla/mux"
)


func GetMeds(w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage)

	query := fmt.Sprintf("SELECT id, name, description, id_pharmacy,(SELECT COUNT(*)  from pharmacy_sh.meds) as count FROM pharmacy_sh.meds LIMIT " + elem + ",10")

	fmt.Println(query)

	var meds []*model.MedSql

	var page model.Page

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id *nullsql.JsonNullInt64
		var name *nullsql.JsonNullString
		var description *nullsql.JsonNullString
		var pharmacyID *nullsql.JsonNullInt64
		var count int
		err = selDB.Scan(&id, &name, &description, &pharmacyID, &count)
		if err != nil {
			panic(err.Error())
		}

		med := model.MedSql{
			ID:   id,
			Name: name,
			Description:  description,
			PharmacyID: pharmacyID,
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

	query := fmt.Sprintf("SELECT * FROM pharmacy_sh.meds WHERE id = "+ nID + "")

	fmt.Println(query)

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	med := model.MedSql{}
	for selDB.Next() {
		var id, pharmacyId *nullsql.JsonNullInt64
		var name, description *nullsql.JsonNullString
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

	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`meds` (`name`, `description`, `id_pharmacy`) VALUES ('%s', '%s','%d' )", med.Name , med.Description, med.PharmacyID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
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

	var query string = "UPDATE `pharmacy_sh`.`meds` SET"

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

	update, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeleteMed(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`meds` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
