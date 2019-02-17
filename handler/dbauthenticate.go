package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user model.LoginUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var response model.RequestResponse
	var pharmacy model.Pharmacy

	rows, err := db.DB.Table("pharmacy").Select("id, cif, address, phone_number, schedule, name, guard, mail").Where("mail = ? and password = ?", user.Mail, user.Password).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, phoneNumber int
			var name, address, scheduler, cif, mail, guard string
			rows.Scan(&id, &cif, &address, &phoneNumber, &scheduler, &name, &guard, &mail)

			pharmacyAux := model.Pharmacy{
				ID:          id,
				Cif:         name,
				Address:     address,
				PhoneNumber: phoneNumber,
				Schedule:    scheduler,
				Name:        name,
				Guard:       guard,
				Mail:        mail,
			}
			pharmacy = pharmacyAux
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_USER)
	}

	loginResponse := model.LoginResponse{
		Pharmacy: &pharmacy,
		Response: response,
	}

	output, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func CheckMail(w http.ResponseWriter, r *http.Request) {

	mail := r.URL.Query().Get("mail")

	var response model.RequestResponse
	var checkMail model.CheckMail

	rows, err := db.DB.Table("pharmacy").Select("id, mail, password").Where("mail = ? ", mail).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id int
			var mail, password string

			rows.Scan(&id, &mail, &password)

			checkMailAux := model.CheckMail{
				ID:   id,
				Mail: mail,
			}

			if password != "" {
				checkMailAux.State = true
			} else {
				checkMailAux.State = false
			}
			checkMail = checkMailAux

		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_USER)
	}

	checkMailResponse := model.CheckMailResponse{
		CheckMail: &checkMail,
		Response:  response,
	}

	output, err := json.Marshal(checkMailResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}
