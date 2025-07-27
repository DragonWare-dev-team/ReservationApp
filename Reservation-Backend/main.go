package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//handlers

func itsrunninhhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, nigga!")
}

func getReservationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reservation)
}

func addReservstionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "This is the post route nigga", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecode(r.body).Decode(&newReservation)

	if err != nil {
		http.Error(w, "Invalid json data", http.StatusBadRequest)
		return
	}

	reservation = append(reservation, newReservation)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newReservation)
}

// vars
type Reservation struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

var reservation = []Reservation{{Name: "Nigga", Time: "69PM"}}

// main func
func main() {
	fmt.Println("Starting server on localhost:8080")
	http.HandleFunc("/hello", itsrunninhhandler)
	http.HandleFunc("/reservation", getReservationsHandler)
	http.ListenAndServe(":8080", nil)
}
