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

func GetTreatmentsCustom(nID string, w http.ResponseWriter, r *http.Request) ([]*model.TreatmentCustomResponse, model.RequestResponse) {

	var treatmentsCustomResponse []*model.TreatmentCustomResponse
	var response model.RequestResponse
	//rows, err  := db.DB.Raw("SELECT id, id_med, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, time, alarm, start_treatment, end_treatment, period FROM pharmacy_sh.treatment_custom WHERE id_user = " + nID +"").Rows()
	rows, err := db.DB.Table("treatment_custom").Select("treatment_custom.id, treatment_custom.id_med, med.name, treatment_custom.time, treatment_custom.alarm, treatment_custom.start_treatment, treatment_custom.end_treatment, treatment_custom.period ").Joins("INNER JOIN med ON med.id =  treatment_custom.id_med").Where("id_user = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, idMed, period int
			var name, time, startTreatment, endTreatment, alarm string
			rows.Scan(&id, &idMed, &name, &time, &alarm, &startTreatment, &endTreatment, &period)

			treatmentsCustom := model.TreatmentCustomResponse{
				ID:             id,
				IDMed:          idMed,
				Name:           name,
				Time:           time,
				Alarm:          alarm,
				StartTreatment: startTreatment,
				EndTreatment:   endTreatment,
				Period:         period,
			}

			treatmentsCustomResponse = append(treatmentsCustomResponse, &treatmentsCustom)

		}

		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_TREATMENTCUSTOM)
	}

	return treatmentsCustomResponse, response
}

func CreateTreatmentCustom(w http.ResponseWriter, r *http.Request) {

	var treatmentCustom model.TreatmentCustom
	var response model.RequestResponse

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &treatmentCustom)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("treatment_custom").Create(&treatmentCustom)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, util.TITLE_TREATMENTCUSTOM)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}

func UpdateTreatmentCustom(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var treatmentCustom model.TreatmentCustom

	err = json.Unmarshal(b, &treatmentCustom)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("treatment_custom").Where("id = ?", nID).Updates(&treatmentCustom)

	util.CheckErr(db.Error)

	if err != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, util.TITLE_TREATMENTCUSTOM)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func DeleteTreatmentCustom(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("treatment_custom").Where("id= ?", nID).Delete(&model.TreatmentCustom{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleNotExistError(int(db.RowsAffected), error.DELETE, util.TITLE_TREATMENTCUSTOM)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}
