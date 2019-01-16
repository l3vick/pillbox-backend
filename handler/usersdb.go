package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage) 

	query := fmt.Sprintf("SELECT id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city, (SELECT COUNT(*)  from pharmacy_sh.users) as count FROM users LIMIT " + elem + ",10 ")

	fmt.Println(query)

	var users []*model.UserNotNull

	var page model.Page

	selDB, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy, count int
		var age, phone_number util.JsonNullInt64
		var name, surname string
		var familyname, address, gender, mail, zip, province, city util.JsonNullString
		err = selDB.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city, &count)

		if err != nil {
			panic(err.Error())
		}

		user := model.UserNotNull{
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

		var index int
		if (count > 10){

			if (count % 10 == 0){
				index = 1
			}else{
				index = 0
			}

			if intPage == 0 {
				page.First = 0
				page.Previous = 0
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			} else if intPage == (count/10) - index {
				page.First = 0
				page.Previous = intPage -1
				page.Next = intPage
				page.Last = (count/10) - index
				page.Count = count
			} else {
				page.First = 0
				page.Previous = intPage-1
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			}
		} else {
			page.First = 0
			page.Previous = 0
			page.Next = 0
			page.Last = 0
			page.Count = 0
		}

	}

	response := model.UserResponseNotNull{
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

	var users []*model.UserNotNull
	var page model.Page

	selDB, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy, count int
		var age, phone_number util.JsonNullInt64
		var name, surname string
		var familyname, address, gender, mail, zip, province, city util.JsonNullString
		err = selDB.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &count, &zip, &province, &city)

		if err != nil {
			panic(err.Error())
		}

		user := model.UserNotNull{
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

		var index int
		if (count > 10){

			if (count % 10 == 0){
				index = 1
			}else{
				index = 0
			}

			if intPage == 0 {
				page.First = 0
				page.Previous = 0
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			} else if intPage == (count/10) - index {
				page.First = 0
				page.Previous = intPage -1
				page.Next = intPage
				page.Last = (count/10) - index
				page.Count = count
			} else {
				page.First = 0
				page.Previous = intPage-1
				page.Next = intPage+1
				page.Last = (count/10) - index
				page.Count = count
			}
		} else {
			page.First = 0
			page.Previous = 0
			page.Next = 0
			page.Last = 0
			page.Count = 0
		}

	}

	response := model.UserResponseNotNull{
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
	nID := r.URL.Query().Get("id")
	user := model.UserNotNull{}

	selDB, err := dbConnector.Query("SELECT * FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy int
		var age, phone_number util.JsonNullInt64
		var name, surname string
		var familyname, address, gender, mail, zip, province, city util.JsonNullString
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

	userJSON, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		// handle error
	}

	w.Write([]byte(userJSON))
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