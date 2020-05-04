package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Testing 123
type Person struct {
	Name string
	Age  int
}

func personCreate(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p Person
	vars := mux.Vars(r)

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	fmt.Println("start")
	fmt.Println("key" + vars["key"])

	//var keys map[string]string

	/*type parms struct {
		key   string
		value string
		pType int16 // 1 string 2 float 3 multi string
	}

	p1 := parms{
	  	key:"pagename",
	  	pType: 1,
	  }

	  keys["pageName"] = ""
	  keys["events"]=""
	  for i; {
	  	key
	  }*/

	keys1 := r.URL.Query().Get("pageName")
	keys2 := r.URL.Query().Get("events")
	//keys, ok := r.URL.Query()["pageName"]
	//keys2, ok2 := r.URL.Query()["Events"]

	//if !ok || len(keys[0]) < 1 {
	//		log.Println("Url Param 'pageName' is missing")
	//		return
	//	}
	//if !ok2 || len(keys2[0]) < 1 {
	//	log.Println("Url Param 'pageName' is missing")
	//		return
	//	}
	fmt.Println("pageName ", keys1, keys2)
	//fmt.Println("pageName 2 ", keys[1])
	//fmt.Println("events ", keys2[0])

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Call")
	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", p)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person/create/{key}", personCreate).Methods("GET")

	http.Handle("/", r)
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}
