package handler

import (
	"net/http"
	"strconv"
	"fmt"
	"encoding/json"
	"io/ioutil"

	"github.com/l3vick/go-pharmacy/model"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(pageNumber)

	elementsPage := intPage * 10

	elem := strconv.Itoa(elementsPage) 

	query := fmt.Sprintf("SELECT id, name, surname, surname_last, med_breakfast, med_launch, med_dinner, alarm_breakfast, alarm_launch, alarm_dinner, age, address, phone, gender, mail, id_pharmacy, (SELECT COUNT(*)  from rds_pharmacy.users) as count FROM users LIMIT " + elem + ",10 ")

	fmt.Println(query)

	var users []*model.User

	var page model.Page

	selDB, err := dbConnector.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy, count, age, phone int
		var name, medBreakfast, medLaunch, medDinner, alarmBreakfast, alarmLaunch, alarmDinner, password, surname, surname_last, address, gender, mail string
		err = selDB.Scan(&id, &name, &surname, &surname_last, &medBreakfast, &medLaunch, &medDinner, &alarmBreakfast, &alarmLaunch, &alarmDinner, &age, &address, &phone, &gender, &mail, &idPharmacy, &count)

		if err != nil {
			panic(err.Error())
		}

		user := model.User{
			ID:             id,
			Name:           name,
			Surname:        surname,
			SurnameLast:    surname_last,
			MedBreakfast:   medBreakfast,
			MedLaunch:      medLaunch,
			MedDinner:      medDinner,
			AlarmBreakfast: alarmBreakfast,
			AlarmLaunch:    alarmLaunch,
			AlarmDinner:    alarmDinner,
			Password:       password,
			Age:			age,
			Address:		address,
			Phone:			phone,
			Gender:			gender,
			Mail:			mail,
			IDPharmacy:     idPharmacy,
		}

		users = append(users, &user)

		var index int
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

	}
	response := model.UserResponse{
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

	query := fmt.Sprintf("SELECT id, name, surname, surname_last, med_breakfast, med_launch, med_dinner, alarm_breakfast, alarm_launch, alarm_dinner, password, age, address, phone, gender, mail, id_pharmacy, (SELECT COUNT(*)  from rds_pharmacy.users where id_pharmacy = " + nID + ") FROM rds_pharmacy.users where id_pharmacy = " + nID + " limit " + elem + ", 10")

	fmt.Println(query)

	var users []*model.User

	var page model.Page

	selDB, err := dbConnector.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, age, phone, id_pharmacy, count int
		var name, surname, surname_last, med_breakfast, med_launch, med_dinner, alarm_breakfast, alarm_launch, alarm_dinner, password, address, gender, mail string
		err = selDB.Scan(&id, &name, &surname, &surname_last, &med_breakfast, &med_launch, &med_dinner, &alarm_breakfast, &alarm_launch, &alarm_dinner, &password, &age, &address, &phone, &gender, &mail, &id_pharmacy, &count )
		if err != nil {
			panic(err.Error())
		}

		user := model.User{
			ID:             id,
			Name:           name,
			Surname: 		surname,
			SurnameLast: 	surname_last,
			MedBreakfast:   med_breakfast,
			MedLaunch:      med_launch,
			MedDinner:      med_dinner,
			AlarmBreakfast: alarm_breakfast,
			AlarmLaunch:    alarm_launch,
			AlarmDinner:    alarm_dinner,
			Password:		password,
			Age:			age,
			Address:		address,
			Phone:			phone,
			Gender:			gender,
			Mail:			mail,
			IDPharmacy: 	id_pharmacy,
		}

		users = append(users, &user)



		var index int
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

	}

	response := model.UserResponse{
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
	user := model.User{}

	selDB, err := dbConnector.Query("SELECT * FROM users WHERE id=?", nID)

	if err != nil {
		panic(err.Error())
	}

	for selDB.Next() {
		var id, idPharmacy, age, phone int
		var name, medBreakfast, medLaunch, medDinner, alarmBreakfast, alarmLaunch, alarmDinner, password, surname, surname_last, address, gender, mail string
		err = selDB.Scan(&id, &name, &surname, &surname_last, &medBreakfast, &medDinner, &medLaunch, &alarmDinner, &alarmLaunch, &alarmBreakfast, &password, &age, &address, &phone, &gender, &mail, &idPharmacy)

		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.Name = name
		user.Surname = surname
		user.SurnameLast = surname_last
		user.MedBreakfast = medBreakfast
		user.MedLaunch = medLaunch
		user.MedDinner = medDinner
		user.AlarmBreakfast = alarmBreakfast
		user.AlarmLaunch = alarmLaunch
		user.AlarmDinner = alarmDinner
		user.Age = age
		user.Address = address
		user.Phone = phone
		user.Gender = gender
		user.Mail = mail
		user.IDPharmacy = idPharmacy
		user.Password = password
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

	query := fmt.Sprintf("INSERT INTO `rds_pharmacy`.`users` (`name`, `surname`, `surname_last`, `med_breakfast`, `med_launch`, `med_dinner`, `alarm_breakfast`, `alarm_launch`, `alarm_dinner`, `password`, `age`, `address`, `phone`, `gender`, `mail, `id_pharmacy`)  VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%s', '%d', '%s', '%s', '%d')", user.Name, user.Surname, user.SurnameLast, user.MedBreakfast, user.MedLaunch, user.MedDinner, user.AlarmBreakfast, user.AlarmLaunch, user.AlarmDinner, user.Password, user.Age, user.Address, user.Phone, user.Gender, user.Mail, user.IDPharmacy)

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

	var query string = fmt.Sprintf("UPDATE `rds_pharmacy`.`users` SET `name` = '%s', `surname` = '%s', `surname_last` = '%s', `med_breakfast` = '%s', `med_launch` = '%s', `med_dinner` = '%s', `alarm_breakfast` = '%s', `alarm_launch` = '%s', `alarm_dinner` = '%s', `age` = '%d', `address` = '%s', `phone` = '%d', `gender` = '%s', `mail` = '%s', `id_pharmacy` = '%d' WHERE (`id` = '%s')", user.Name, user.Surname, user.SurnameLast, user.MedBreakfast, user.MedLaunch, user.MedDinner, user.AlarmBreakfast, user.AlarmBreakfast, user.AlarmBreakfast, user.Age, user.Address, user.Phone, user.Gender, user.Mail, user.IDPharmacy, nID)

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

	query := fmt.Sprintf("DELETE FROM `rds_pharmacy`.`users` WHERE (`id` = '%s')", nID)

	fmt.Println(query)
	insert, err := dbConnector.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}