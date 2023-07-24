// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person - Our struct for all persons
type Profession struct {
	Id   int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var Professions []Profession

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllProfessions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllProfessions")
	json.NewEncoder(w).Encode(Professions)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/professions", returnAllProfessions)
	log.Fatal(http.ListenAndServe(":4881", myRouter))
}

func main() {
	Professions = []Profession{
		{Id: 1, Name: "Kesselflicker", Desc: "KFL"},
		{Id: 2, Name: "Pinselmacher", Desc: "PMR"},
		{Id: 3, Name: "Hahnenspanner", Desc: "HHSP"},
	}
	handleRequests()
}
