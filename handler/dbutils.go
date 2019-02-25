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

func GetScheduleByUserID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	var treatmentsResponse model.TreatmentsResponse
	var response model.RequestResponse
	var responseCustom model.RequestResponse
	var responseTiming model.RequestResponse

	treatmentsResponse.Treatments, response = GetTreatmentsByUserID(nID, w, r)
	treatmentsResponse.Response = append(treatmentsResponse.Response, response)
	treatmentsResponse.TreatmentsCustom, responseCustom = GetTreatmentsCustomByUserID(nID, w, r)
	treatmentsResponse.Response = append(treatmentsResponse.Response, responseCustom)
	treatmentsResponse.Timing, responseTiming = GetTimingByUserID(nID, w, r)
	treatmentsResponse.Response = append(treatmentsResponse.Response, responseTiming)

	output, err := json.Marshal(treatmentsResponse)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user model.User

	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("user").Where("id = ?", nID).Updates(&user)

	util.CheckErr(db.Error)

	if err != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, util.TITLE_USER)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}
