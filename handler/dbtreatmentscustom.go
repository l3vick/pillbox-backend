package handler

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/l3vick/go-pharmacy/model"
)

func GetTreatmentsCustom(nID string, w http.ResponseWriter, r *http.Request) ([]*model.TreatmentCustomResponse){

	var treatmentsCustomResponse []*model.TreatmentCustomResponse

	query := fmt.Sprintf("SELECT id, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, time, alarm, start_treatment, end_treatment, period FROM pharmacy_sh.treatment_custom WHERE id_user = " + nID +"")

	fmt.Println(query)

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	for selDB.Next() {
		var id, period int
		var alarm byte
		var name, time, start_treatment, end_treatment string
		err = selDB.Scan(&id, &name, &time, &alarm, &start_treatment, &end_treatment, &period)
		if err != nil {
			panic(err.Error())
		}

		treatmentsCustom := model.TreatmentCustomResponse {
			ID: id,
			Name: name,
			Time: time,
			Alarm: alarm,
			StartTreatment: start_treatment,
			EndTreatment: end_treatment,
			Period: period,
		}

		treatmentsCustomResponse = append(treatmentsCustomResponse, &treatmentsCustom)
	}
	
	if err := selDB.Err(); err != nil {
        panic(err.Error())
	}

	if err := selDB.Close(); err != nil {
		panic(err.Error())
	}
	return treatmentsCustomResponse
}

func CreateTreatmentCustom(w http.ResponseWriter, r *http.Request) {
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
	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`treatment` (`id_user`, `id_med`, `morning`, `afternoon`, `evening`, `end_treatment`)  VALUES('%d', '%d', '%d', '%d', '%d', '%s')", treatment.IDUser, treatment.IDMed, treatment.Morning, treatment.Afternoon, treatment.Evening, treatment.EndTreatment)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func UpdateTreatmentCustom(w http.ResponseWriter, r *http.Request) {

}

func DeleteTreatmentCustom(w http.ResponseWriter, r *http.Request) {

}