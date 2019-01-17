package handler

import (
	"github.com/l3vick/go-pharmacy/nullsql"
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

	query := fmt.Sprintf("SELECT id, cif, address, phone_number, schedule, `name`, guard, mail FROM pharmacy_sh.pharmacy WHERE mail = '%s'  and password = '%s'", user.Mail, user.Password)

	fmt.Println(query)
	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	pharmacy := model.PharmacySql{}
	for selDB.Next() {
		var id nullsql.JsonNullInt64
		var numberPhone, guard *nullsql.JsonNullInt64
		var name, address, scheduler, cif, mail *nullsql.JsonNullString
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

	}

	output, err := json.Marshal(pharmacy)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
	defer selDB.Close()
}