package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
	"github.com/l3vick/go-pharmacy/error"
)

const title string = "Treatment"
var response model.RequestResponse

func GetTreatmentsByUserID(w http.ResponseWriter, r *http.Request) {

	var treatmentsResponse model.TreatmentsResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	treatmentsResponse = GetTreatments(nID, w, r)
	treatmentsResponse.TreatmentsCustom = GetTreatmentsCustom(nID, w, r)
	treatmentsResponse.Timing = GetTiming(nID, w, r)

	output, err := json.Marshal(treatmentsResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func GetTreatments(nID string, w http.ResponseWriter, r *http.Request) (model.TreatmentsResponse) {

	var treatmentResponse model.TreatmentsResponse

	query := fmt.Sprintf("SELECT id, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, morning, afternoon, evening, start_treatment, end_treatment FROM pharmacy_sh.treatment WHERE id_user = " + nID +"")

	fmt.Println(query)

	selct, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var mornings []*model.Morning
	var afternoons []*model.Afternoon
	var evenings []*model.Evening

	for selct.Next() {
		var id int
		var morning, afternoon, evening byte
		var name, start_treatment, end_treatment string

		err = selct.Scan(&id, &name, &morning, &afternoon, &evening, &start_treatment, &end_treatment)
		
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

	query := "INSERT INTO pharmacy_sh.treatment (id_user, id_med, morning, afternoon, evening, start_treatment, end_treatment) VALUES( ?, ?, ?, ?, ?, ?, ?)"

	insert, err := dbConnector.Prepare(query)

	util.CheckErr(err)

	insertExec, err := insert.Exec(treatment.IDUser, treatment.IDMed,  util.BoolToByte(treatment.Morning),  util.BoolToByte(treatment.Afternoon),  util.BoolToByte(treatment.Evening), treatment.StartTreatment, treatment.EndTreatment)

	util.CheckErr(err)

	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		response = error.HandleEmptyRowsError(insertExec, error.Insert, title)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func UpdateTreatment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var treatment model.Treatment

	err = json.Unmarshal(b, &treatment)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	query := "UPDATE pharmacy_sh.treatment SET id_med = ?, morning = ?, afternoon = ?, evening = ?, start_treatment = ?, end_treatment = ? WHERE (id = ?)"

	update, err := dbConnector.Prepare(query)

	util.CheckErr(err)

	updateExec, err := update.Exec(treatment.IDMed,  util.BoolToByte(treatment.Morning),  util.BoolToByte(treatment.Afternoon),  util.BoolToByte(treatment.Evening), treatment.StartTreatment, treatment.EndTreatment, nID)

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		response = error.HandleEmptyRowsError(updateExec, error.Update, title)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func DeleteTreatment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	nID := vars["id"]

	query := "DELETE FROM pharmacy_sh.treatment WHERE (id = ?)"

	delete, err := dbConnector.Prepare(query)

	util.CheckErr(err)

	deleteExec, err := delete.Exec(nID)

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		response = error.HandleEmptyRowsError(deleteExec, error.Delete, title)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}
