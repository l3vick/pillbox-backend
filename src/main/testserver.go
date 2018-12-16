package main

import (
	"database/sql"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Estoy funcionando!!!!-------------------" + message
	w.Write([]byte(message))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "rds_pharmacy_00:phar00macy@tcp(rdspharmacy00.ctiytnyzqbi7.us-east-2.rds.amazonaws.com)/rdspharmacy00")
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Get Userv !!!-db:  error" + err.Error()
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	http.HandleFunc("/hola", getUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
