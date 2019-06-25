package main

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id int
	animalType string
	nickname string
	zone int
	age int
}

func main (){
	//connecting to the database
	db, err := sql.Open("mysql","root:mysqlpass@/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//general query with arguments
	rows, err := db.Query("select * from dino.animals where  age > ?", 10) // the ques mark is the arg which we pass later in the func
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		 a :=   animal{}
		 err := rows.Scan(&a.id,&a.animalType,&a.nickname,&a.zone,&a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	// Err returns the error, if any, that was encountered during iteration.
	// Err may be called after an explicit or implicit Close.

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

	// query s single row even if there re morethn one results give only one row
	row := db.QueryRow("select * from dino.animals where id > ?", 2 )
	a := animal{}
	err = row.Scan(&a.id,&a.animalType,&a.nickname,&a.zone,&a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	////insert a row
	//result, err := db.Exec("Insert into dino.animals (animals_type,nick,zone,age) values ('Carnotaurus', 'Carno', 3, 22)")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())


	// update a row
	//result, err := db.Exec("Update dino.animals set age = ? where id =?",16,2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())
	//

}