package main

import (
	_ "errors"
	"github.com/l3vick/go-pharmacy/db"
	"github.com/l3vick/go-pharmacy/handler"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "App Farmacias" + message
	w.Write([]byte(message))
}

func main() {
	db.ConectDB()

	r := mux.NewRouter()
	r.HandleFunc("/", root).Methods("GET")

	r.HandleFunc("/meds", handler.GetMeds).Methods("GET")
	r.HandleFunc("/meds/{id}", handler.GetMed).Methods("GET")
	r.HandleFunc("/meds", handler.CreateMed).Methods("POST")
	r.HandleFunc("/meds/{id}", handler.UpdateMed).Methods("PUT")
	r.HandleFunc("/meds/{id}", handler.DeleteMed).Methods("DELETE")

	r.HandleFunc("/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/resetpassword/{id}", handler.ResetPassword).Methods("PUT")
	r.HandleFunc("/filter/{filter}", handler.FilterUser).Methods("GET")

	r.HandleFunc("/pharmacies", handler.GetPharmacies).Methods("GET")
	r.HandleFunc("/pharmacies/{id}/users", handler.GetUsersByPharmacyID).Methods("GET")
	r.HandleFunc("/pharmacies/{id}", handler.GetPharmacy).Methods("GET")
	r.HandleFunc("/pharmacies", handler.CreatePharmacy).Methods("POST")
	r.HandleFunc("/pharmacies/{id}", handler.UpdatePharmacy).Methods("PUT")
	r.HandleFunc("/pharmacies/{id}", handler.DeletePharmacy).Methods("DELETE")

	r.HandleFunc("/treatments/{id}", handler.GetTreatmentsByUserID).Methods("GET")
	r.HandleFunc("/treatments", handler.CreateTreatment).Methods("POST")
	r.HandleFunc("/treatments/{id}", handler.UpdateTreatment).Methods("PUT")
	r.HandleFunc("/treatments/{id}", handler.DeleteTreatment).Methods("DELETE")

	r.HandleFunc("/treatments/custom", handler.CreateTreatmentCustom).Methods("POST")
	r.HandleFunc("/treatments/custom/{id}", handler.UpdateTreatmentCustom).Methods("PUT")
	r.HandleFunc("/treatments/custom/{id}", handler.DeleteTreatmentCustom).Methods("DELETE")

	r.HandleFunc("/timing", handler.CreateTiming).Methods("POST")
	r.HandleFunc("/timing/{id}", handler.UpdateTiming).Methods("PUT")
	r.HandleFunc("/timing/{id}", handler.DeleteTiming).Methods("DELETE")

	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.HandleFunc("/checkMail", handler.CheckMail).Methods("GET")

	http.Handle("/", &MyServer{r})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if req.Method == "OPTIONS" {
		return
	}

	s.r.ServeHTTP(w, req)
}
