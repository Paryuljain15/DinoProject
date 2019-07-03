package dinoapi

import (
	"Practice/DinoProject/databaselayer"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type DinoRESTAPIHandler struct{
	dbhandler databaselayer.DinoDBHandler
}


func newDinoRESTAPIHandler(db databaselayer.DinoDBHandler) *DinoRESTAPIHandler {
	return &DinoRESTAPIHandler{
		dbhandler:db,
	}
}

func (handler *DinoRESTAPIHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  //variable for map from string to string
	criteria, ok := vars["SearchCriteria"]
	if !ok {  //if we don't find the search criteria
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex , or
								to search by type via /api/dinos/type/velociraptor`)
		return
	}
	searchkey, ok := vars["search"]
	if !ok {  //if we don't find the search criteria
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `No search criteria found, you can either search by nickname via /api/dinos/nickname/rex , or
								to search by type via /api/dinos/type/velociraptor`)
		return
	}
	var animal databaselayer.Animal
	var animals []databaselayer.Animal
	var err error
	switch strings.ToLower(criteria) {
	case "nickname" :
		animal, err = handler.dbhandler.GetDynoByNickname(searchkey)
	case "type":
		animals, err = handler.dbhandler.GetDynosByType(searchkey)
		if len(animals) > 0 {
			json.NewEncoder(w).Encode(animals)
			return
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error occured while querying animals %v", err)
	}


	json.NewEncoder(w).Encode(animal)
}

func (handler *DinoRESTAPIHandler) editsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  //variable for map from string to string
	operation, ok := vars["Operation"]
	if !ok {  //if we don't find the search criteria
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `Operation was not found, you can either use /api/dinos/add to add a new animal, or
								/api/dinos/edit/rex to edit an animal with nickname rex`)
		return
	}
	var animal databaselayer.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Could not decode the request body to json %v", err)
		return
	}
	switch strings.ToLower(operation) {
	case "add" :
		err = handler.dbhandler.AddAnimal(animal)
	case "edit":
		//api/dinos/edit/rex relative url
		nickname := r.RequestURI[len("/api/dinos/edit/"):]  // cutting the front path to get the nick
		log.Println("edit requested for nickname", nickname)
		err = handler.dbhandler.UpdateAnimal(animal, nickname)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error occured while processing request %v", err)
	}
}
