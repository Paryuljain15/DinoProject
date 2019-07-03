package dinoapi

import (
	"Practice/DinoProject/databaselayer"
	"github.com/gorilla/mux"
	"net/http"
)

//dino API
//HTTp GET for search /api/dinos/nickname/rex  or by /api/dinos/type/velociraptor  (these are all relative urls)
//HTTP POST to add or edit /api/dinos/add or (example localhost:8181)/api/dinos/edit
// router is something which binds a location to some action like get, post
//if editing dis done on basis on nickname (for eg.) then
//HTTP POST to /api/dinos/nickname/rex, rex will be edited using the data obtained from the http json body of the request

// The HTTP POST request comes with a json body which hosys the data tp be added or used to the edit

func RunApi(endpoint string, db databaselayer.DinoDBHandler) error {
	//endpoint example : localhost:8080
	r := mux.NewRouter()
	RunAPIOnRouter(r, db)
	return http.ListenAndServe(endpoint, r)
}

func RunAPIOnRouter(r *mux.Router,db databaselayer.DinoDBHandler) {
	handler := newDinoRESTAPIHandler(db)

	apirouter := r.PathPrefix("/api/dinos").Subrouter() // this subrouter has the given relative url till api/dinos and will act as a prefix for further locations

	apirouter.Methods("GET").PathPrefix("/{SearchCriteria}/{search}").HandlerFunc(handler.searchHandler)  //curly braces means it is flexible
	apirouter.Methods("POST").PathPrefix("/{Operation}").HandlerFunc(handler.editsHandler)			// if we had used Path function then we would have to define the exact location
}