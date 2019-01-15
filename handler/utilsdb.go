package handler

import (
	"net/http"
	"fmt" 
	"encoding/json"
	"io/ioutil"

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

	query := fmt.Sprintf("SELECT * FROM pharmacy_sh.pharmacy WHERE phone_number =  %d  and password = '%s'", user.Phone, user.Password)

	fmt.Println(query)
	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	pharmacy := model.Pharmacy{}
	for selDB.Next() {
		var id, phoneNumber, guard int
		var cif, street, schedule, name, password, mail string
		err = selDB.Scan(&id, &cif, &street, &phoneNumber, &schedule, &name, &guard, &password, &mail)

		pharmacy.ID = id
		pharmacy.Cif = cif
		pharmacy.Street = street
		pharmacy.NumberPhone = phoneNumber
		pharmacy.Schedule = schedule
		pharmacy.Name = name
		pharmacy.Guard = guard
		pharmacy.Password = password
		pharmacy.Mail = mail


		if err != nil {
			panic(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}

	}

	output, err := json.Marshal(pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
	defer selDB.Close()
}