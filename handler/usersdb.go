package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/nullsql"
	"github.com/l3vick/go-pharmacy/util"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage) 

	query := fmt.Sprintf("SELECT id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city, (SELECT COUNT(*)  from pharmacy_sh.users) as count FROM users LIMIT " + elem + ",10 ")

	fmt.Println(query)

	var users []*model.UserSql

	var page model.Page

	selDB, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy nullsql.JsonNullInt64
		var count int
		var age, phone_number *nullsql.JsonNullInt64
		var name, surname, familyname, address, gender, mail, zip, province, city *nullsql.JsonNullString
		err = selDB.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city, &count)

		if err != nil {
			panic(err.Error())
		}

		user := model.UserSql{
			ID:             id,
			Name:           name,
			SurName:       	surname,
			FamilyName:    	familyname,
			Age:			age,
			Address:		address,
			Phone:			phone_number,
			Gender:			gender,
			Mail:			mail,
			IDPharmacy:     idPharmacy,
			Zip:			zip,
			Province:		province,
			City:			city,
		}

		users = append(users, &user)

		page = util.GetPage(count, intPage)
	}

	response := model.UserResponseSql{
		Users: users,
		Page: page,
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)	
}

func GetUsersByPharmacyID(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	nID := vars["id"]
	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	query := fmt.Sprintf("SELECT id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city, (SELECT COUNT(*)  from pharmacy_sh.users where id_pharmacy = " + nID + ") FROM pharmacy_sh.users where id_pharmacy = " + nID + " limit " + elem + ", 10")

	fmt.Println(query)

	var users []*model.UserSql
	var page model.Page

	selDB, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy nullsql.JsonNullInt64
		var age, phone_number *nullsql.JsonNullInt64
		var count int
		var name, surname, familyname, address, gender, mail, zip, province, city *nullsql.JsonNullString
		err = selDB.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &count, &zip, &province, &city)

		if err != nil {
			panic(err.Error())
		}

		user := model.UserSql{
			ID:             id,
			Name:           name,
			SurName: 		surname,
			FamilyName: 	familyname,
			Age:			age,
			Address:		address,
			Phone:			phone_number,
			Gender:			gender,
			Mail:			mail,
			IDPharmacy: 	idPharmacy,
			Zip:			zip,
			Province:		province,
			City:			city,

		}

		users = append(users, &user)

		page = util.GetPage(count, intPage)
	}

	response := model.UserResponseSql{
		Users: users,
		Page: page,
	}

	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]
	user := model.UserSql{}

	selDB, err := dbConnector.Query("SELECT id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy nullsql.JsonNullInt64
		var age, phone_number *nullsql.JsonNullInt64
		var name, surname, familyname, address, gender, mail, zip, province, city *nullsql.JsonNullString
		err = selDB.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.SurName = surname
		user.FamilyName = familyname
		user.Age = age
		user.Address = address
		user.Phone = phone_number
		user.Gender = gender
		user.Mail = mail
		user.IDPharmacy = idPharmacy
		user.Zip = zip
		user.Province = province
		user.City = city
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(output)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)

	query := fmt.Sprintf("INSERT INTO `pharmacy_sh`.`users` (`name`, `surname`, `familyname`, `password`, `age`, `address`, `phone_number`, `gender`, `mail`, `zip`, `province`, `city`, `id_pharmacy`)  VALUES('%s', '%s', '%s', '%s', '%d', '%s', '%d', '%s', '%s', '%s', '%s', '%s', '%d')", user.Name, user.SurName, user.FamilyName, user.Password, user.Age, user.Address, user.Phone, user.Gender, user.Mail, user.Zip, user.Province, user.City, user.IDPharmacy)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(output)

	var query string = fmt.Sprintf("UPDATE `pharmacy_sh`.`users` SET `name` = '%s', `surname` = '%s', `familyname` = '%s', `age` = '%d', `address` = '%s', `phone_number` = '%d', `gender` = '%s', `mail` = '%s', `id_pharmacy` = '%d', `zip` = '%s', `province` = '%s', `city` = '%s' WHERE (`id` = '%s')", user.Name, user.SurName, user.FamilyName, user.Age, user.Address, user.Phone, user.Gender, user.Mail, user.IDPharmacy, user.Zip, user.Province, user.City, nID)

	fmt.Println(query)
	update, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	nID := vars["id"]

	query := fmt.Sprintf("DELETE FROM `pharmacy_sh`.`users` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}