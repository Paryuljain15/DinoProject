package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Animal struct {
	ID 			int
	AnimalType	string
	Nickname	string
	Zone 		int
	Age 		int
}
func main() {
	//Carnotaurus, Carno,3,22
	data := &Animal{
		AnimalType: "Velociraptor",
		Nickname:   "patro",
		Zone:       3,
		Age:        14,    // age changed to check edit
	}
	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	//for adding an animal

	//resp, err := http.Post("http://localhost:8181/api/dinos/add","application/json", &b)  // we here used a byte buffer because it implements io.writer whereas data var does not
	//if err != nil || resp.Status != "200 OK" {  // or resp.StatusCode != 200 would also work
	//	log.Fatal(err)
//     }

	// for editing


	resp, err := http.Post("http://localhost:8181/api/dinos/edit/patro", "apllication/json", &b)
	if err != nil || resp.Status != "200 OK" {  // or resp.StatusCode != 200 would also work
		log.Fatal(err)
	}
}


