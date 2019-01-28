package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
)

func GetTreatmentsByUserID(w http.ResponseWriter, r *http.Request) {

	var treatmentsResponse model.TreatmentsResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	treatmentsResponse = GetTreatment(nID, w, r)
	treatmentsResponse.TreatmentsCustom = GetTreatmentsCustom(nID, w, r)
	treatmentsResponse.Timing = GetTiming(nID, w, r)

	output, err := json.Marshal(treatmentsResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func GetTreatment(nID string, w http.ResponseWriter, r *http.Request) (model.TreatmentsResponse) {

	var treatmentResponse model.TreatmentsResponse

	query := fmt.Sprintf("SELECT id, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, morning, afternoon, evening, start_treatment, end_treatment FROM pharmacy_sh.treatment WHERE id_user = " + nID +"")

	fmt.Println(query)

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	var mornings []*model.Morning
	var afternoons []*model.Afternoon
	var evenings []*model.Evening

	for selDB.Next() {
		var id int
		var morning, afternoon, evening byte
		var name, start_treatment, end_treatment string

		err = selDB.Scan(&id, &name, &morning, &afternoon, &evening, &start_treatment, &end_treatment)
		if err != nil {
			panic(err.Error())
		}

		if morning == 1 {
			morningAux := model.Morning {
				ID: id,
				Name: name,
				StartTreatment: start_treatment,
				EndTreatment: end_treatment,
			}
			mornings = append(mornings, &morningAux)
		}

		if afternoon == 1 {
			afternoonAux := model.Afternoon {
				ID: id,
				Name: name,
				StartTreatment: start_treatment,
				EndTreatment: end_treatment,
			}
			afternoons = append(afternoons, &afternoonAux)
		}

		if evening == 1 {
			eveningAux := model.Evening {
				ID: id,
				Name: name,
				StartTreatment: start_treatment,
				EndTreatment: end_treatment,
			}
			evenings = append(evenings, &eveningAux)
		}
	}

	if err := selDB.Err(); err != nil {
		panic(err.Error())
	}

	if err := selDB.Close(); err != nil {
		panic(err.Error())
	}

	treatmentResponse.Morning = mornings
	treatmentResponse.Afternoon = afternoons
	treatmentResponse.Evening = evenings

	return treatmentResponse
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
	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`treatment` (`id_user`, `id_med`, `morning`, `afternoon`, `evening`, `end_treatment`)  VALUES('%d', '%d', '%d', '%d', '%d', '%s')", treatment.IDUser, treatment.IDMed, treatment.Morning, treatment.Afternoon, treatment.Evening, treatment.EndTreatment)

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
