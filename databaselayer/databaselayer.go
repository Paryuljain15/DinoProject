package databaselayer

import "errors"

const (
	MYSQL uint8 = iota
	SQLITE
	POSTGRESQL
	MONGODB
)
type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal,error)
	GetDynoByNickname(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID 				int 		`bson:"-"`
	AnimalType 		string 		`bson:"animal_type"`
	Nickname 		string 		`bson:"nickname"`
	Zone 			int			`bson:"zone"`
	Age   			int			`bson:"age"`
}


var DBTypeNotSupported = errors.New("the Database type provide is not supported") // error string should not be capitalised

//factory function

func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {

	switch dbtype {
	//case MYSQL:
	//	return NewMySQLHandler(connection)
	case MONGODB:
		return NewMongoDBHandler(connection)
	//case SQLITE:
	//	return NewSQLiteHandler(connection)
	case POSTGRESQL :
		return NewPQHandler(connection)

	}
	return nil, DBTypeNotSupported
}