package handler

import (
	/*"encoding/json"
	"fmt"
	"github.com/gorilla/mux"*/

	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"

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

			/*timing.Morning = morning
			timing.Afternoon = afternoon
			timing.Evening = evening
			timing.MorningTime = morningTime
			timing.AfternoonTime = afternoonTime
			timing.EveningTime = evening

			timingResponse = timing
			*/timing := model.Timing{
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

func CreateTiming(w http.ResponseWriter, r *http.Request) {
	/*
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

		output, err := json.Marshal(timing)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(output)

		query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`timing` (`id_user`, `morning`, `afternoon`, `evening`, `morning_time`, `afternoon_time`, `evening_time`)  VALUES('%d', '%d', '%d', '%d', '%s', '%s', '%s')", timing.Id_User, timing.Morning, timing.Afternoon, timing.Evening, timing.Morning_Time, timing.Afternoon_Time, timing.Evening_Time)

		fmt.Println(query)
		insert, err := db.DB.Query(query)

		var timingResponse model.RequestResponse
		if err != nil {
			timingResponse.Code = 500
			timingResponse.Message = err.Error()
		} else {
			timingResponse.Code = 200
			timingResponse.Message = "Timing creado con éxito"
		}

		output, err2 := json.Marshal(timingResponse)
		if err2 != nil {
			http.Error(w, err.Error(), 501)
			return
		}

		w.Write(output)

		defer insert.Close()
	*/
}

func UpdateTiming(w http.ResponseWriter, r *http.Request) {
	/*
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

		var query string = fmt.Sprintf("UPDATE `pharmacy_sh`.`timing` SET `morning` = '%d', `afternoon` = '%d', `evening` = '%d', `morning_time` = '%s', `afternoon_time` = '%s', `evening_time` = '%s' WHERE (`id_user` = '%s')", timing.Morning, timing.Afternoon, timing.Evening, timing.Morning_Time, timing.Afternoon_Time, timing.Evening_Time, nID)

		fmt.Println(query)
		update, err := db.DB.Query(query)

		var timingResponse model.RequestResponse
		if err != nil {
			timingResponse.Code = 500
			timingResponse.Message = err.Error()
		} else {
			timingResponse.Code = 200
			timingResponse.Message = "Timing actualizado con éxito"
		}

		output, err := json.Marshal(timingResponse)
		if err != nil {
			http.Error(w, err.Error(), 501)
			return
		}

		w.Write(output)

		defer update.Close()
	*/
}

func DeleteTiming(w http.ResponseWriter, r *http.Request) {
	/*
		vars := mux.Vars(r)
		nID := vars["id"]

		query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`timing` WHERE (`id_user` = '%s')", nID)

		fmt.Println(query)
		insert, err := db.DB.Query(query)

		var timingResponse model.RequestResponse
		if err != nil {
			timingResponse.Code = 500
			timingResponse.Message = err.Error()
		} else {
			timingResponse.Code = 200
			timingResponse.Message = "Timing borrado con éxito"
		}

		output, err := json.Marshal(timingResponse)
		if err != nil {
			http.Error(w, err.Error(), 501)
			return
		}

		w.Write(output)

		defer insert.Close()
	*/
}
