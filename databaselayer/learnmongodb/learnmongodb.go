package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	//"fmt"
	"log"

	"gopkg.in/mgo.v2"
)
type animal struct {
	// ID int 'bson: "id"'
	AnimalType string `bson:"animal_type"` // uppercase as to get it exported
	Nickname string     `bson :"nickname"`
	Zone int 			`bson : "zone"`
	Age int 			`bson : "age"`
}
 func main() {
	session,err := mgo.Dial("mongodb://localhost:27017")
 	if err != nil {
 		log.Fatal(err)
	}
	defer session.Close()
	//get a collection
 	animalcollection := session.DB("Dino").C("animals") // if the database and collection doesn't exist it wil create one
	/*animals := []interface{}{  // an empty interface so that we can put anything in it
		animal{
			AnimalType:"Tyrannosaurus Rex",
			Nickname:"rex",
			Zone:1,
			Age:12,
			}, animal {
			AnimalType:"Velociraptor",
			Nickname:"rapter",
			Zone:2,
			Age:15,
		}, animal {
			AnimalType:"Velociraptor",
			Nickname:"velo",
			Zone:2,
			Age:9,
		},

	}
	err = animalcollection.Insert(animals...) // unfurling a slice
	if err != nil {
		log.Fatal(err)
	}
	*/

	//Updating a document
	//err = animalcollection.Update(bson.M{"nickname":"rapter"}, bson.M{"$set":bson.M{"age":18}})


	//Removing a document

	//err = animalcollection.Remove(bson.M{"nickname" : "velo"})
	//if err != nil {
	//	log.Fatal(err)
	//}


 	//for general queries
 	// age > 10 and zone in (1,2)
 	query := bson.M{
 		"age": bson.M{"$gt" : 10, },
 		"zone" : bson.M{"$in":[]int{1,2}},
	}
	results := []animal{}
	animalcollection.Find(query).All(&results)  // all says all the results for one reult we could have written one
	fmt.Println(results)
 }
