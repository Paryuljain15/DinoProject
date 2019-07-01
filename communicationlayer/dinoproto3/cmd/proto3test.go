package main

import (
	"Practice/DinoProject/communicationlayer/dinoproto3"
	"fmt"
	"github.com/golang/protobuf/proto"

	"flag"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

/*
	1-We will serialize some data via proto3
	2-We will send this data via TCP t a different service
	3-We will deserialize the data via proto3, and print out the extracted values

	A-A TCP client needs to be written to send the data
	B- A TCP Server to receive the data

we are building a client application that is exchanging data with the server
*/

func main(){
	op :=flag.String("op", "s","s for server, and c for client")    //when this binary runs it should support an op flag which we will have a default value of s and c
	flag.Parse() 				//so proto3test -op s => will run as server. proto3test -op c => will run as a client
	switch strings.ToLower(*op){
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}
}


func RunProto3Server() {
	l,err := net.Listen("tcp",":8282") //will create a tcp sever at 8282 port
	if err != nil {
		log.Fatal(err)
	}
	for i := 0;i < 2;i++ {			//this is an endless for loop but in industrial processes it will have some conditions as oin this case it will only take 3 requests
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)  // data will be the slice of byte we will obtain on the tcp channel
			if err != nil {
				return
			}
			a := &dinoproto3.Animal{}
			err = proto.Unmarshal(data, a)  //a will reflect protocol serialized data which was embedded in rhe data
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(a)
		}(c)
	}
}

func RunProto3Client() {
	a := &dinoproto3.Animal{
		Id:			2,				//proto.Int(1),  this are written for comparison with proto2
		AnimalType: "Trex",		//proto.String("Raptor"),
		Nickname :  "rex", 				//proto.String("rapto"),
		Zone:		3,			//proto.Int(1),
		Age:		12,				//proto.Int(15),
	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "localhost:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)
}