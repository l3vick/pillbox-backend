package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func CreateTreatment(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

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
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, util.TITLE_TREATMENT)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}

func GetTreatment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	var treatmentResponse model.TreatmentResponse
	var response model.RequestResponse

	rows, err := db.DB.Table("treatment").Select("treatment.id, med.name, treatment.morning, treatment.afternoon, treatment.evening, treatment.start_treatment, treatment.end_treatment").Joins("INNER JOIN med ON med.id =  treatment.id_med").Where("treatment.id = ?", nID).Rows()

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

			treatmentAux := model.TreatmentName{
				ID:             id,
				Name:           name,
				Morning:        morning,
				Afternoon:      afternoon,
				Evening:        evening,
				StartTreatment: start_treatment,
				EndTreatment:   end_treatment,
			}
			treatmentResponse.Treatment = treatmentAux
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_TREATMENT)
	}

	treatmentResponse.Response = response

	output, err := json.Marshal(treatmentResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func UpdateTreatment(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

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
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, util.TITLE_TREATMENT)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

}

func DeleteTreatment(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("treatment").Where("id= ?", nID).Delete(&model.Treatment{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleNotExistError(int(db.RowsAffected), error.DELETE, util.TITLE_TREATMENT)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)

}

func GetTreatmentsByUserID(nID string, w http.ResponseWriter, r *http.Request) (model.Treatments, model.RequestResponse) {

	var treatment model.Treatments
	var response model.RequestResponse
	var morningTreatments []*model.TreatmentName
	var afternoonTreatments []*model.TreatmentName
	var eveningTreatments []*model.TreatmentName

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
				treatmentAux := model.TreatmentName{
					ID:             id,
					Name:           name,
					Morning:        morning,
					Afternoon:      afternoon,
					Evening:        evening,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				morningTreatments = append(morningTreatments, &treatmentAux)
			}

			if afternoon == "true" {
				treatmentAux := model.TreatmentName{
					ID:             id,
					Name:           name,
					Morning:        morning,
					Afternoon:      afternoon,
					Evening:        evening,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				afternoonTreatments = append(afternoonTreatments, &treatmentAux)
			}

			if evening == "true" {
				treatmentAux := model.TreatmentName{
					ID:             id,
					Name:           name,
					Morning:        morning,
					Afternoon:      afternoon,
					Evening:        evening,
					StartTreatment: start_treatment,
					EndTreatment:   end_treatment,
				}
				eveningTreatments = append(eveningTreatments, &treatmentAux)
			}
		}
		treatment.TreatmentsMorning = morningTreatments
		treatment.TreatmentsAfternoon = afternoonTreatments
		treatment.TreatmentsEvening = eveningTreatments
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_TREATMENT)
	}

	return treatment, response
}
