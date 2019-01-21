package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// State struct (Model)
type State struct {
	Type       string `json:"type"`
	Icon       string `json:"icon"`
	Title      string `json:"title"`
	Value      string `json:"value"`
	FooterText string `json:"footerText"`
	FooterIcon string `json:"footerIcon"`
}

// User struct (Model)
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// Init States var as a slice State struct
var states []State

// enable CORS
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Get all States
func getStates(w http.ResponseWriter, r *http.Request) {
	states[0].Value = strconv.Itoa(random(10, 200)) + "GB"
	json.NewEncoder(w).Encode(states)
}

//authenticate User
func authenticate(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.Name == "john" && user.Password == "password" {
		json.NewEncoder(w).Encode(user)
	}
}

//generate random number between range
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()
	// Hardcoded data
	states = append(states, State{Type: "warning", Icon: "ti-server", Title: "Capacity", FooterText: "Updated now", FooterIcon: "ti-reload"})
	states = append(states, State{Type: "success", Icon: "ti-wallet", Title: "Revenue", Value: "$1,345", FooterText: "Last day", FooterIcon: "ti-calendar"})
	states = append(states, State{Type: "danger", Icon: "ti-pulse", Title: "Errors", Value: "23", FooterText: "In the last hour", FooterIcon: "ti-timer"})
	states = append(states, State{Type: "info", Icon: "ti-twitter-alt", Title: "Followers", Value: "+45", FooterText: "Updated now", FooterIcon: "ti-reload"})

	// Route handles & endpoints
	r.HandleFunc("/states", getStates).Methods("GET")
	r.HandleFunc("/authenticate", authenticate).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
