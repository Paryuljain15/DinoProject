package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	)

type MongoDBHandler struct {
	*mgo.Session
}
func NewMongodbHandler(connection string)(*MongoDBHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBHandler{
		Session : s,
	}, err
}
func (handler *MongoDBHandler) GetAvailableDynos() ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(nil).All(&animals)
	return animals, err
}
func (handler *MongoDBHandler) GetDynoByNickname(nickname string) (Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	a := Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"nickname" : nickname}).One(&a)
	return a, err
}
func (handler *MongoDBHandler) GetDynosByType(dinotype string) ([]Animal, error) {
	s := handler.getFreshSession()
	defer s.Close()
	animals := []Animal{}
	err := s.DB("Dino").C("animals").Find(bson.M{"animal_type":dinotype}).All(&animals)
	return animals, err
}
func (handler *MongoDBHandler) AddAnimal(a Animal) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Insert(a)
}

func (handler *MongoDBHandler) UpdateAnimal(a Animal, nname string) error {
	s := handler.getFreshSession()
	defer s.Close()
	return s.DB("Dino").C("animals").Update(bson.M{"nickname" : nname}, a)
}

func (handler *MongoDBHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}