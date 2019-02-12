package handler

import (
	/*"encoding/json"
	"fmt"

	"github.com/gorilla/mux"*/

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
	"io/ioutil"

	/*	"github.com/l3vick/go-pharmacy/util"
		"github.com/l3vick/go-pharmacy/db"
		"io/ioutil"*/
	"net/http"
)

const TITLE_TIMING string = "Timing"

func GetTiming(idUser string, w http.ResponseWriter, r *http.Request) (model.Timing, model.RequestResponse) {

	var response model.RequestResponse
	var timingResponse model.Timing
	//rows, err  := db.DB.Raw("SELECT id, id_med, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, time, alarm, start_treatment, end_treatment, period FROM pharmacy_sh.treatment_custom WHERE id_user = " + nID +"").Rows()
	rows, err := db.DB.Table("timing").Select("morning, afternoon, evening, morning_time, afternoon_time, evening_time").Where("id_user = ?", idUser).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			//var timing *model.Timing
			count = count + 1
			var morningTime, afternoonTime, eveningTime, morning, afternoon, evening string

			rows.Scan(&morning, &afternoon, &evening, &morningTime, &afternoonTime, &eveningTime)

			timing := model.Timing{
				Morning:       morning,
				Afternoon:     afternoon,
				Evening:       evening,
				MorningTime:   morningTime,
				AfternoonTime: afternoonTime,
				EveningTime:   eveningTime,
			}
			timingResponse = timing
		}
		response = error.HandleNoRowsError(count, error.SELECT, TITLE_TIMING)
	}

	return timingResponse, response

}

func GetTimingByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	var response model.RequestResponse
	var timing model.Timing
	//rows, err  := db.DB.Raw("SELECT id, id_med, (SELECT name FROM pharmacy_sh.med WHERE id = id_med) as name, time, alarm, start_treatment, end_treatment, period FROM pharmacy_sh.treatment_custom WHERE id_user = " + nID +"").Rows()
	rows, err := db.DB.Table("timing").Select("morning, afternoon, evening, morning_time, afternoon_time, evening_time").Where("id_user = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			//var timing *model.Timing
			count = count + 1
			var morningTime, afternoonTime, eveningTime, morning, afternoon, evening string

			rows.Scan(&morning, &afternoon, &evening, &morningTime, &afternoonTime, &eveningTime)

			timingAux := model.Timing{
				Morning:       morning,
				Afternoon:     afternoon,
				Evening:       evening,
				MorningTime:   morningTime,
				AfternoonTime: afternoonTime,
				EveningTime:   eveningTime,
			}
			timing = timingAux
		}
		response = error.HandleNoRowsError(count, error.SELECT, TITLE_TIMING)
	}

	timingResponse := model.TimingResponse{
		Timing:   timing,
		Response: response,
	}

	output, err := json.Marshal(timingResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

}

func CreateTiming(w http.ResponseWriter, r *http.Request) {

	var timing model.Timing
	var response model.RequestResponse

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &timing)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("timing").Create(&timing)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, TITLE_TIMING)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

}

func UpdateTiming(w http.ResponseWriter, r *http.Request) {

	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var timing model.Timing

	err = json.Unmarshal(b, &timing)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("timing").Where("id_user = ?", nID).Updates(&timing)

	util.CheckErr(db.Error)

	if err != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.Update, TITLE_TIMING)
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func DeleteTiming(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("timing").Where("id_user= ?", nID).Delete(&model.Timing{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.DELETE, TITLE_TIMING)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}
