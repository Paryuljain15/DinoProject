package databaselayer

import (
	"database/sql"
	"fmt"
	"log"
)

type SQLHandler struct {

	*sql.DB
}
func (handler *SQLHandler) GetAvailableDynos() ([]Animal, error) {
	return handler.sendQuery("Select * from Animals")
}


func (handler *SQLHandler) GetDynoByNickname(nickname string) (Animal, error) {
	row := handler.QueryRow(fmt.Sprintf("Select * from Animals where nickname = '%s'",nickname))
	a := Animal{}
	err := row.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
	return a, err
}
func (handler *SQLHandler) GetDynosByType (dinotype string) ([]Animal, error) {
	return handler.sendQuery(fmt.Sprintf("Select * from Animals where AnimalType = '%s", dinotype))
}
func (handler *SQLHandler) AddAnimal(a Animal) error {
	_, err := handler.Exec(fmt.Sprintf("Insert into Animals (animal_type, nickname, zone, age) values ('%s', '%s',%d, %d)",a.AnimalType, a.Nickname, a.Zone, a.Age))
	return err
}
func (handler *SQLHandler) UpdateAnimal(a Animal, nname string) error {
	_, err := handler.Exec(fmt.Sprintf("Update Animals set animal_type = '%s', nickname = '%s', zone = %d, age = %d where nickname = '%s'", a.AnimalType, a.Nickname, a.Zone, a.Age, nname))
	return err
}
func (handler *SQLHandler) sendQuery(q string) ([]Animal, error) {
	Animals := []Animal{}
	rows, err := handler.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.AnimalType, &a.Nickname, &a.Zone, &a.Age)
		if err != nil {
			log.Println(err)
			continue
		}
		Animals = append(Animals, a)

	}
	return Animals, rows.Err()
}
