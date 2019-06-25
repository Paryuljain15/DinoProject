package main

import (
	"fmt"
	//"fmt"
	"log"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type animal struct {
	gorm.Model  // it iniatates id and prints created time,deleted time
	//ID	int `gorm:"primary_key;not null;unique;AUTO_INCREMENT"`
	Animal_type string `gorm:"type:TEXT"`
	Nickname string `gorm:"type:TEXT"`
	Zone	int 	`gorm:"type:INTEGER"`
	Age 	int 	`gorm:"type:INTEGER"`

}
func main() {
	db, err := gorm.Open("mysql","root:mysqlpass@/dino" )
	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	// there are only some basic funcs

	db.DropTableIfExists(&animal{}) // will delete the table  which has animal struct
	db.AutoMigrate(&animal{}) // will add any missing fields to table which has this object, to name the table it will add 's'
	db.Table("dinos").DropTableIfExists(&animal{})								// to the name so the table name will become animals
	// to create a table
	db.Table("dinos").CreateTable(&animal{})

	a := animal{
		Animal_type:"Tyrannosaurus Rex",
		Nickname:"rex",
		Zone:1,
		Age:12,
	}
	db.Save(&a)  // vs create()
	db.Table("dinos").Create(&a) // to explicitly add data to a row
	a = animal {
		Animal_type:"Velicoraptor",
		Nickname:"Velo",
		Zone:2,
		Age:15,
	}
	db.Save(&a)

	// updates the database
	db.Table("animals").Where("nickname = ? and zone =?","Velo",2).Update("age",16)

	//queries
	var  b animal
	animals := []animal{}
	db.Find(&animals,"age > ?",12) // finds in both tables unless specified
	db.Debug().Table("dinos").Find(&b,"age > ?",10) // debug shows the query whuch was run
	fmt.Println(b)
	animals = append(animals,b)
	fmt.Println(animals)




}



