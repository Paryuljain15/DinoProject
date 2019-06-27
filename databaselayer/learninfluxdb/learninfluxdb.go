package main

import (
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := queryDB(c,"dino", `Select * from weightmeasures where "animal_type" = 'Tyrannosaurus Rex'`)
	if err != nil {
		log.Fatal(err)
	}
	for _, v :=  range res {
		log.Println("messages: ", v.Messages)
		for _, s := range v.Series {
				log.Println("series name: ", s.Name)
				log.Println("series columns: ", s.Columns)
				log.Println("series values: ", s.Values)
		}

	}

}
func queryDB(cl client.Client, database,cmd string) (res []client.Result, err error) {
q := client.Query{
Command:cmd,
Database:database,
}
if response, err := cl.Query(q); err == nil {
if response.Error() !=nil {
return res, response.Error()
}
res = response.Results

} else {
return res, err  // both empty
}
return res, nil
}