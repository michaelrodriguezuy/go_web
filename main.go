package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID       uint64 `json:"id"`
	FirstName     string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

var users []User 

func init() {
	users =[]User{
		{
			ID:       1,	
			FirstName:     "Juan",
			LastName: "Pérez",
			Email: "m@m.com",
			Age:      30,
		},
		{
			ID:       2,
			FirstName:     "Pedro",
			LastName: "Gómez",
			Email: "d@d.com",
			Age:      25,
		},
	}
}


func main() {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		UserServer(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		
		GetAllUsers(w)
		//fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "impeca el Get")
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "impeca el Post")
	case http.MethodPut:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "impeca el Put")
	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)
		//codigo 204
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusNoContent, "impeca el Delete")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		//codigo 405
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusMethodNotAllowed, "No existe el metodo")
	}
}



func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, data any) { //en este caso users es una entidad del tipo interfaz(any en las ultimas versiones), la idea con esto es que le pueda mandar cualquier entidad y la tome

	value, _:= json.Marshal(data) //marshall convierte una estructura a un json

	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status": %d, "data": "%s"}`, status, value)

	// if err := json.NewEncoder(w).Encode(data); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
