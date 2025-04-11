package main

import (
	"fmt"
	"log"
	"net/http"
)

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
		w.WriteHeader(http.StatusOK)		
		fmt.Fprintf(w, `{"status": %d, "message": "%s"}`, http.StatusOK, "impeca el Get") 
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
