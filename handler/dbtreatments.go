package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

const TITLE_TREATMENT string = "Treatment"

var response model.RequestResponse

func GetAllTreatmentsByUserID(w http.ResponseWriter, r *http.Request) {

	var treatmentsResponse model.TreatmentsResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	treatmentsResponse = GetTreatmentsByUserID(nID, w, r)

	output, err := json.Marshal(treatmentsResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func GetTreatmentsByUserID(nID string, w http.ResponseWriter, r *http.Request) model.TreatmentsResponse {

	var treatmentResponse model.TreatmentsResponse
	var mornings []*model.Morning
	var afternoons []*model.Afternoon
	var evenings []*model.Evening
	var responseSecondary model.RequestResponse

	//rows, err  := db.DB.Raw("SELECT id, (SELECT name FROM pharmacy_sh.med WHERE id = id_med), morning, afternoon, evening, start_treatment, end_treatment FROM pharmacy_sh.treatment WHERE id_user = " + nID +"").Rows()
	rows, err := db.DB.Table("treatment").Select("treatment.id, med.name, treatment.morning, treatment.afternoon, treatment.evening, treatment.start_treatment, treatment.end_treatment").Joins("INNER JOIN med ON med.id =  treatment.id_med").Where("id_user = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id int
			var name, start_treatment, end_treatment, morning, afternoon, evening string

			rows.Scan(&id, &name, &morning, &afternoon, &evening, &start_treatment, &end_treatment)

			if morning == "true" {
				morningAux := model.Morning{
					ID:             id,
					Name:           name,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				mornings = append(mornings, &morningAux)
			}

			if afternoon == "true" {
				afternoonAux := model.Afternoon{
					ID:             id,
					Name:           name,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				afternoons = append(afternoons, &afternoonAux)
			}

			if evening == "true" {
				eveningAux := model.Evening{
					ID:             id,
					Name:           name,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				evenings = append(evenings, &eveningAux)
			}
		}
		response = error.HandleNoRowsError(count, error.Select, TITLE_TREATMENT)
		treatmentResponse.Response = append(treatmentResponse.Response, response)
		fmt.Println(fmt.Sprintf(" code: %d  message: %s ", response.Code, response.Message))
	}

	treatmentResponse.Morning = mornings
	treatmentResponse.Afternoon = afternoons
	treatmentResponse.Evening = evenings
	treatmentResponse.TreatmentsCustom, responseSecondary = GetTreatmentsCustom(nID, w, r)
	treatmentResponse.Timing = GetTiming(nID, w, r)
	fmt.Println("--------------------------------")
	fmt.Println(fmt.Sprintf(" code: %d  message: %s ", response.Code, response.Message))
	treatmentResponse.Response = append(treatmentResponse.Response, responseSecondary)

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

	db := db.DB.Table("treatment").Create(&treatment)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Insert, TITLE_TREATMENT)
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

	db := db.DB.Table("treatment").Where("id = ?", nID).Updates(treatment)

	util.CheckErr(db.Error)

	if err != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, TITLE_TREATMENT)
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

	db := db.DB.Table("treatment").Where("id= ?", nID).Delete(&model.Treatment{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Delete, TITLE_TREATMENT)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)

}
