package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
type animal struct {
	id int
	animalType string
	nickname string
	zone int
	age int
}
func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgresqlpass dbname = dino sslmode = disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//general query with args.
	rows, err := db.Query("select * from animals where age > $1",5) // unlike ? here we use $1 to pass args
	handlerows(rows,err)
	// query a single row
	row := db.QueryRow("SELECT * From animals where age > $1",5)
	a := animal{}
	err = row.Scan(&a.id, &a.zone,&a.animalType, &a.nickname,  &a.age)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(a)
	//insert a row

	//result,err := db.Exec("Insert into animals (zone,animal_type,nickname,age) values ($1,'Carnotaurus','carno',$2)",3,20)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(result.LastInsertId()) // not available for postgre
	//fmt.Println(result.RowsAffected())


	//update a row
	//res, err := db.Exec("Update animals set age =$1 where id =$2",16,2)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.LastInsertId()) // not available for postgre
	//fmt.Println(res.RowsAffected())


	//returning an id which was changed
	//var id int
	//db.QueryRow("update animals set age =$1 where id = $2 returning id",17,2).Scan(&id)
	//fmt.Println(id)


	//prepare queries to use multiple times, this also improves perfomance
// usable for all drivers
	fmt.Println("Statements ... ")
	stmt, err :=  ("Select * from animals where age >$1") // no args are there
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//for age >5
	rows,err = stmt.Query(5)
	handlerows(rows,err)
	//foe age > 10
	rows, err = stmt.Query(10)
	handlerows(rows,err)

 testTransactions(db)
}
func handlerows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id,&a.zone, &a.animalType, &a.nickname,  &a.age)
		if err != nil {
			log.Println(err)
			continue 
		}
		animals = append(animals, a)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)
	}

func testTransactions(db *sql.DB){
	fmt.Println("Transactions... ")
	tx,err := db.Begin()
	if err != nil {
		log.Fatal(err)
		}
	defer tx.Rollback() // reverses all changes made before commit

	stmt, err := db.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(15)
	handlerows(rows,err)
	rows, err = stmt.Query(17)
	handlerows(rows,err)

	err = tx.Commit()  //will commit all transactions you have done to databse
	if err != nil {
		log.Fatal(err)
	}
}
