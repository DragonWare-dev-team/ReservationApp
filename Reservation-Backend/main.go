package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func itsrunninhhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, nigga!")
}

func getReservationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(reservation)
}

type Reservation struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

var reservation = []Reservation{{Name: "Nigga", Time: "69PM"}}

func main() {
	fmt.Println("Starting server on localhost:8080")
	http.HandleFunc("/hello", itsrunninhhandler)
	http.HandleFunc("/reservation", getReservationsHandler)
	http.ListenAndServe(":8080", nil)
}
