package handler

import (
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
	"github.com/gorilla/mux"
)


func GetMeds(w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage)

	query := fmt.Sprintf("SELECT id, name, description, pharmacy_id,(SELECT COUNT(*)  from pharmacy_sh.meds) as count FROM pharmacy_sh.meds LIMIT " + elem + ",10")

	fmt.Println(query)

	var meds []*model.Med

	var page model.Page

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id int
		var name string
		var description string
		var pharmacyID util.JsonNullInt64
		var count int
		err = selDB.Scan(&id, &name, &description, &pharmacyID, &count)
		if err != nil {
			panic(err.Error())
		}

		med := model.Med{
			ID:   id,
			Name: name,
			Description:  description,
			PharmacyID: pharmacyID,
		}
		meds = append(meds, &med)
		

		var index int
		if (count > 10){

			if (count % 10 == 0){
				index = 1
			}else{
				index = 0
			}

			if intPage == 0 {
				page.First = 0
				page.Previous = 0
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			} else if intPage == (count/10) - index {
				page.First = 0
				page.Previous = intPage -1
				page.Next = intPage
				page.Last = (count/10) - index
				page.Count = count
			} else {
				page.First = 0
				page.Previous = intPage-1
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			}
		} else {
			page.First = 0
			page.Previous = 0
			page.Next = 0
			page.Last = 0
			page.Count = 0
		}
		

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

	selDB, err := dbConnector.Query("SELECT * FROM pharmacy_sh.meds WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	med := model.Med{}
	for selDB.Next() {
		var id int
		var pharmacyId util.JsonNullInt64
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

	var med model.MedInt
	err = json.Unmarshal(b, &med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`meds` (`name`, `description`, `pharmacy_id`) VALUES ('%s', '%s','%d' )", med.Name, med.Description, med.PharmacyID)

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

	var query string = fmt.Sprintf("UPDATE `pharmacy_sh`.`meds` SET `name` = '%s', `description` = '%s' WHERE (`id` = '%s')", med.Name, med.Description, nID)

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
