package handler

import (
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/gorilla/mux"
)


func GetMeds(w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage) 

	query := fmt.Sprintf("SELECT id, name, pvp, (SELECT COUNT(*)  from rds_pharmacy.med) as count FROM med LIMIT " + elem + ",10")

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
		var pvp int
		var count int
		err = selDB.Scan(&id, &name, &pvp, &count)
		if err != nil {
			panic(err.Error())
		}
		med := model.Med{
			ID:   id,
			Name: name,
			Pvp:  pvp,
		}
		meds = append(meds, &med)

		var index int
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

	selDB, err := dbConnector.Query("SELECT * FROM med WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}

	med := model.Med{}
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

	fmt.Println(med.Name)
	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)

	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`med` (`name`, `pvp`) VALUES('%s','%d')", med.Name, med.Pvp)

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

	output, err := json.Marshal(med)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`med` SET `name` = '%s', `pvp` = '%d' WHERE (`id` = '%s')", med.Name, med.Pvp, nID)

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

	query := fmt.Sprintf("DELETE FROM `rds_pharmacy`.`med` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}
