
//This file needs improvement as it panics dur=e to accessing unbuilt index

package main

import (
	"fmt"
	"github.com/influxdata/influxdb1-client/v2"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

//collect weights of each animal on a frequent basis => time series dataset
// this data gets stored in influxdb, so that we can use it later


var animaltags = []string{"Tyrannosaurus Rex;rex","Velociraptor:rapto","Velociraptor:velo","Carnotasaurus:carno"}

const myDB = "dino"

func main() {
	cl, err := client.NewHTTPClient(client.HTTPConfig{
		Addr :"http://localhost:8086",
		Username :"",
		Password:"",
	})
	if err!= nil {
		log.Fatal(err)
	}
	_,err = queryDB(cl,"","Create DATABASE " +myDB) // creating dino database, database name in second argument is empty because we are actually creating one
	if err != nil {
		log.Fatal(err)
	}
	// create a batch points object
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:myDB,
		Precision:"s", // seconds
	})
	if err!= nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	detectSignal := checkStopOSSignals(&wg)
	rand.Seed(time.Now().UnixNano())  //
//this loop generates random data for weights
	for !(*detectSignal) {
		animaltag := animaltags[rand.Intn(len(animaltags))]
		split := strings.Split(animaltag,";")
		tags := map[string]string{
			"animal_type" : split[0],
			"nickname" : split[1],
		}
		fields := map[string]interface{}{
			"weight" : rand.Intn(300) +1,
		}
		fmt.Println(animaltag, fields["weight"])
		pt, err := client.NewPoint("weightmeasures", tags, fields, time.Now())
		if err != nil {
			log.Println(err)
			continue
		}
		bp.AddPoint(pt)
		time.Sleep(1*time.Second)
	}
	log.Println("Exiting signal triggered, writing data... ")
	if err := cl.Write(bp); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	log.Println("Exiting program...")
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
func checkStopOSSignals(wg *sync.WaitGroup) *bool {
	Signal := false
	go func(s *bool) {
		wg.Add(1)
		ch := make(chan os.Signal)  //make a channel of type os.signal
		signal.Notify(ch, syscall.SIGINT,syscall.SIGTERM)  // notify to ch channel if siginterruption or sig termination signal received
		<-ch  //wait on the channel until something is received
		log.Println("Exit signals received...")
		*s = true  // changing Signal to true
		wg.Done()
	}(&Signal)
	return &Signal
}