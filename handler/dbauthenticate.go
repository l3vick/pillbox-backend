package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/l3vick/go-pharmacy/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user model.UserLogin
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	query := fmt.Sprintf("SELECT id, cif, address, phone_number, schedule, `name`, guard, mail FROM pharmacy_sh.pharmacy WHERE mail = '%s' and password = '%s'", user.Mail, user.Password)

	fmt.Println(query)
	selDB, err := dbConnector.Query(query)

	var checkMailResponse model.RequestResponse
	if err != nil {
		checkMailResponse.Code = 500
		checkMailResponse.Message = err.Error()
		output, _ := json.Marshal(checkMailResponse)
		w.Write(output)
	} else {
		pharmacy := model.Pharmacy{}
		for selDB.Next() {
			var id, numberPhone, guard int
			var name, address, scheduler, cif, mail string
			err = selDB.Scan(&id, &cif, &address, &numberPhone, &scheduler, &name, &guard, &mail)

			pharmacy.ID = id
			pharmacy.Cif = cif
			pharmacy.Address = address
			pharmacy.NumberPhone = numberPhone
			pharmacy.Schedule = scheduler
			pharmacy.Name = name
			pharmacy.Guard = guard
			pharmacy.Mail = mail

			if err != nil {
				panic(err.Error())
				http.Error(w, err.Error(), 500)
				return
			}

			output, err := json.Marshal(pharmacy)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write(output)
		}
	}
	defer selDB.Close()
}

func CheckMail(w http.ResponseWriter, r *http.Request) {
	mailRequest := r.URL.Query().Get("mail")

	query := fmt.Sprintf("SELECT id, mail, password FROM pharmacy_sh.pharmacy WHERE mail = '%s'", mailRequest)
	fmt.Println(query)

	selDB, err := dbConnector.Query(query)

	var checkMailResponse model.RequestResponse
	if err != nil {
		checkMailResponse.Code = 500
		checkMailResponse.Message = err.Error()
		output, _ := json.Marshal(checkMailResponse)
		w.Write(output)
	} else {
		pharmacy := model.PharmacyR{}
		for selDB.Next() {
			var id int
			var mail, password string
			err = selDB.Scan(&id, &mail, &password)

			if err != nil {
				panic(err.Error())
				http.Error(w, err.Error(), 500)
				return
			}

			pharmacy.ID = id
			pharmacy.Mail = mail
			if password == "" {
				pharmacy.State = false
			} else {
				pharmacy.State = true
			}

			output, err := json.Marshal(pharmacy)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write(output)
		}
	}
	defer selDB.Close()
}
