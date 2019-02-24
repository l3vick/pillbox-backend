package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/error"
	"github.com/l3vick/go-pharmacy/model"
	"github.com/l3vick/go-pharmacy/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var response model.RequestResponse

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	db := db.DB.Table("user").Create(&user)

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleEmptyRowsError(db.RowsAffected, error.INSERT, util.TITLE_USER)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	var response model.RequestResponse
	var user model.User

	rows, err := db.DB.Table("user").Select("id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city").Where("id = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, idPharmacy, age, phone_number int
			var name, surname, familyname, address, gender, mail, zip, province, city string
			rows.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city)

			userAux := model.User{
				ID:          id,
				Name:        name,
				Surname:     surname,
				Familyname:  familyname,
				Age:         age,
				Address:     address,
				PhoneNumber: phone_number,
				Gender:      gender,
				Mail:        mail,
				IDPharmacy:  idPharmacy,
				Zip:         zip,
				Province:    province,
				City:        city,
			}
			user = userAux
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_USER)
	}

	userResponse := model.UserResponse{
		User:     &user,
		Response: response,
	}

	output, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var response model.RequestResponse

	vars := mux.Vars(r)

	nID := vars["id"]

	db := db.DB.Table("user").Where("id = ?", nID).Delete(&model.User{})

	util.CheckErr(db.Error)

	if db.Error != nil {
		response = error.HandleMysqlError(db.Error)
	} else {
		response = error.HandleNotExistError(int(db.RowsAffected), error.DELETE, util.TITLE_USER)
	}

	output, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}
	w.Write(output)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name != "" {
		FilterUserByName(name, w, r)
	} else {
		GetAllUsers(w, r)
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var users []*model.User
	var page model.Page
	var count int

	rows, err := db.DB.Table("user").Select("id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city, (SELECT COUNT(*)  from pharmacy_sh.user) as count").Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		for rows.Next() {
			var id, idPharmacy, age, phone_number int
			var name, surname, familyname, address, gender, mail, zip, province, city string
			rows.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city, &count)

			user := model.User{
				ID:          id,
				Name:        name,
				Surname:     surname,
				Familyname:  familyname,
				Age:         age,
				Address:     address,
				PhoneNumber: phone_number,
				Gender:      gender,
				Mail:        mail,
				IDPharmacy:  idPharmacy,
				Zip:         zip,
				Province:    province,
				City:        city,
			}
			users = append(users, &user)
		}
		page = util.GetPage(count, intPage)
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_USER)
	}

	userResponse := model.UsersResponse{
		Users:    users,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)

}

func FilterUserByName(name string, w http.ResponseWriter, r *http.Request) {

	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var users []*model.User
	var page model.Page

	rows, err := db.DB.Table("user").Select("id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city").Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Offset(elem).Limit(10).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		var count = 0
		for rows.Next() {
			count = count + 1
			var id, idPharmacy, age, phone_number int
			var name, surname, familyname, address, gender, mail, zip, province, city string
			rows.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city)

			userAux := model.User{
				ID:          id,
				Name:        name,
				Surname:     surname,
				Familyname:  familyname,
				Age:         age,
				Address:     address,
				PhoneNumber: phone_number,
				Gender:      gender,
				Mail:        mail,
				IDPharmacy:  idPharmacy,
				Zip:         zip,
				Province:    province,
				City:        city,
			}
			users = append(users, &userAux)
		}
		response = error.HandleNotExistError(count, error.SELECT, util.TITLE_USER)
		page = util.GetPage(count, intPage)
	}

	usersResponse := model.UsersResponse{
		Users:    users,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(usersResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}

func GetUsersByPharmacyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nID := vars["id"]

	pageNumber := r.URL.Query().Get("page")
	intPage, err := strconv.Atoi(pageNumber)
	elementsPage := intPage * 10
	elem := strconv.Itoa(elementsPage)

	var response model.RequestResponse
	var users []*model.User
	var page model.Page
	var count int

	rows, err := db.DB.Table("user").Select("id, name, surname, familyname, age, address, phone_number, gender, mail, id_pharmacy, zip, province, city, (SELECT COUNT(*)  from pharmacy_sh.user WHERE id_pharmacy  = "+nID+") as count").Offset(elem).Limit(10).Where("id_pharmacy = ?", nID).Rows()

	defer rows.Close()

	util.CheckErr(err)

	if err != nil {
		response = error.HandleMysqlError(err)
	} else {
		for rows.Next() {
			var id, idPharmacy, age, phone_number int
			var name, surname, familyname, address, gender, mail, zip, province, city string
			rows.Scan(&id, &name, &surname, &familyname, &age, &address, &phone_number, &gender, &mail, &idPharmacy, &zip, &province, &city, &count)

			user := model.User{
				ID:          id,
				Name:        name,
				Surname:     surname,
				Familyname:  familyname,
				Age:         age,
				Address:     address,
				PhoneNumber: phone_number,
				Gender:      gender,
				Mail:        mail,
				IDPharmacy:  idPharmacy,
				Zip:         zip,
				Province:    province,
				City:        city,
			}

			users = append(users, &user)
		}
		page = util.GetPage(count, intPage)
		response = error.HandleNoRowsError(count, error.SELECT, util.TITLE_USER)
	}

	userResponse := model.UsersResponse{
		Users:    users,
		Page:     page,
		Response: response,
	}

	output, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, err.Error(), 501)
		return
	}

	w.Write(output)
}
