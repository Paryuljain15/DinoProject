package main

import (
	"../dinowebportal"
	"encoding/json"
	"log"
	"os"
	)
import _ "github.com/go-sql-driver/mysql"


type configuration struct {
	ServerAddress string `json:"webserver"` // so that this file knows how is this connected in json
}
func main() {
	file, err:= os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)  // new is used to create a pointer of passed type
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address",config.ServerAddress)
	dinowebportal.RunWebPortal(config.ServerAddress)
}