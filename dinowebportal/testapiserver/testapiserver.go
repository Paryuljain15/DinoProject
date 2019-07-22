package main

import (
	"Practice/DinoProject/databaselayer"
	"Practice/DinoProject/dinowebportal/dinoapi"
	"fmt"
	"log"
)

func main() {
	db, err := databaselayer.GetDatabaseHandler(databaselayer.MONGODB, "mongodb://127.0.0.1")
	fmt.Println(uint8(databaselayer.MONGODB))
	if err != nil {
		log.Fatal(err)
	}

	dinoapi.RunApi("192.124.120.118:8181", db)

}
