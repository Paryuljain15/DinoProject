package main

import (
	"Practice/DinoProject/communicationlayer/dinoproto2"
	"fmt"
	"github.com/golang/protobuf/proto"

	"flag"
	"io/ioutil"
	"log"
	"net"
	"strings"

)

/*
	1-We will serialize some data via proto2
	2-We will send this data via TCP t a different service
	3-We will deserialize the data via proto2, and print out the extracted values

	A-A TCP client needs to be written to send the data
	B- A TCP Server to receive the data
*/

func main(){
	op :=flag.String("op", "s","s for server, and c for client")    //when this binary runs it should support an op flag which we will have a default value of s and c
	flag.Parse() 				//so proto2test -op s => will run as server. proto2test -op c => will run as a client
	switch strings.ToLower(*op){
	case "s":
		RunProto2Server()
	case "c":
		RunProto2Client()
	}
}


func RunProto2Server() {
	l,err := net.Listen("tcp",":8282") //will create a tcp sever at 8282 port
	if err != nil {
		log.Fatal(err)
	}
	for {			//this is an endless for loop but in industrial processes it will have some conditions
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
			a := &dinoproto2.Animal{}
			proto.Unmarshal(data, a)  //a will reflect protocol serialized data which was embedded in rhe data
			fmt.Println(a)
		}(c)
	}
}

func RunProto2Client() {
	a := &dinoproto2.Animal{
	Id:			proto.Int(3),
	AnimalType: proto.String("Raptor"),
	Nickname :  proto.String("rapto"),
	Zone:		proto.Int(2),
	Age:		proto.Int(12),
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