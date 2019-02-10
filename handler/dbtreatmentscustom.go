package handler

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func GetTreatmentsCustom(nID string, w http.ResponseWriter, r *http.Request) ([]*model.TreatmentCustomResponse){

	var treatmentsCustomResponse []*model.TreatmentCustomResponse

	query := fmt.Sprintf("SELECT id, id_med, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, time, alarm, start_treatment, end_treatment, period FROM pharmacy_sh.treatment_custom WHERE id_user = " + nID +"")

	fmt.Println(query)

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	for selDB.Next() {
		var id, idMed, period int
		var alarm byte
		var name, time, start_treatment, end_treatment string
		err = selDB.Scan(&id, &idMed, &name, &time, &alarm, &start_treatment, &end_treatment, &period)
		if err != nil {
			panic(err.Error())
		}

		treatmentsCustom := model.TreatmentCustomResponse {
			ID: id,
			IDMed: idMed,
			Name: name,
			Time: time,
			Alarm:  util.ByteToBool(alarm),
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
	var treatment model.TreatmentCustom
	err = json.Unmarshal(b, &treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`treatment_custom` (`id_user`, `id_med`, `time`, `alarm`, `start_treatment`, `end_treatment`, `period`)  VALUES('%d', '%d', '%s', '%d', '%s', '%s', '%d')", treatment.IDUser, treatment.IDMed, treatment.Time, util.BoolToByte(treatment.Alarm), treatment.StartTreatment, treatment.EndTreatment, treatment.Period)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	output, err := json.Marshal(treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func UpdateTreatmentCustom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var treatment model.TreatmentCustom

	err = json.Unmarshal(b, &treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var query  = fmt.Sprintf("UPDATE `pharmacy_sh`.`treatment_custom` SET `id_med` = '%d', `time` = '%s', `alarm` = '%d', `start_treatment` = '%s', `end_treatment` = '%s', `period` = '%d' WHERE (`id` = '%s')", treatment.IDMed, treatment.Time, util.BoolToByte(treatment.Alarm), treatment.StartTreatment, treatment.EndTreatment, treatment.Period, nID)

	fmt.Println(query)

	update, err := dbConnector.Query(query)

	var response model.RequestResponse
	if err != nil {
		response.Code = 500
		response.Message = err.Error()
	} else {
		response.Code = 200
		response.Message = "Treatment custom actualizado con éxito"
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer update.Close()
}

func DeleteTreatmentCustom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`treatment_custom` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)

	var response model.RequestResponse
	if err != nil {
		response.Code = 500
		response.Message = err.Error()
	} else {
		response.Code = 200
		response.Message = "Treatment custom borrado con éxito"
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

	defer insert.Close()
}