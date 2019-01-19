package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
)

func GetTreatments(w http.ResponseWriter, r *http.Request) {

}

func GetTreatment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	nID := vars["id"]

	query := fmt.Sprintf("SELECT * FROM pharmacy_sh.treatment WHERE id = " + nID + "")

	fmt.Println(query)

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	treatment := model.Treatment{}
	for selDB.Next() {
		var id, idUser, idMed int
		var morning, afternoon, evening bool
		var endTreatment string
		err = selDB.Scan(&id, &idUser, &idMed, &morning, &afternoon, &evening, &endTreatment)
		if err != nil {
			panic(err.Error())
		}
		treatment.ID = id
		treatment.IDUser = idUser
		treatment.IDMed = idMed
		treatment.Morning = morning
		treatment.Afternoon = afternoon
		treatment.Evening = evening
		treatment.EndTreatment = endTreatment
	}

	output, err := json.Marshal(treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(output)
}

func CreateTreatment(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal
	var treatment model.Treatment
	err = json.Unmarshal(b, &treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`treatments` (`id_user`, `id_med`, `morning`, `afternoon`, `evening`, `end_treatment`)  VALUES('%d', '%d', '%t', '%t', '%t', '%s')", treatment.IDUser, treatment.IDMed, treatment.Morning, treatment.Afternoon, treatment.Evening, treatment.EndTreatment)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateTreatment(w http.ResponseWriter, r *http.Request) {

}

func DeleteTreatment(w http.ResponseWriter, r *http.Request) {

}
