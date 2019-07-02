package main

import (
	"Practice/DinoProject/communicationlayer/dinogrpc"
	"Practice/DinoProject/databaselayer"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	op :=flag.String("op", "s","s for server, and c for client")    //when this binary runs it should support an op flag which we will have a default value of s and c
	flag.Parse() 				//so proto2test -op s => will run as server. proto2test -op c => will run as a client
	switch strings.ToLower(*op){
	case "s":
		runGRPCServer()
	case "c":
		runGRPCClient()
	}

}

func runGRPCServer() {
	//grpclog.Println("Starting GRPC Server")
	fmt.Println("Starting GRPC Server")
	lis, err := net.Listen("tcp", "localhost:8282")
	//grpclog.Infoln("Listening on 127.0.0.1:8282")
	fmt.Println("Listening in 127.0.0.1:8282")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	//var opts []grpc.ServerOption
	grpcServer := grpc.NewServer()    //opts ...  this input isnt necceasary as variadic
	dinoServer, err := dinogrpc.NewDinoGrpcServer(databaselayer.MONGODB, "mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}
	dinogrpc.RegisterDinoServiceServer(grpcServer, dinoServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func runGRPCClient() {
	conn, err := grpc.Dial("localhost:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := dinogrpc.NewDinoServiceClient(conn)
	input := ""
	fmt.Println("All animals? (y/n)")
	fmt.Scanln(&input)
	if strings.EqualFold(input, "y") {  //checks if the arguement strings are equal
		animals, err := client.GetAllAnimals(context.Background(),&dinogrpc.Request{})  //empty context otherwise used for timeouts etc
		if err != nil {
			log.Fatal(err)
		}
		for {
			animal, err := animals.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				grpclog.Fatal(err)
			}
			fmt.Println(animal)
		}
		return
	}
	fmt.Println("Nickname ?")
	fmt.Scanln(&input)
	a, err := client.GetAnimal(context.Background(), &dinogrpc.Request{Nickname: input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*a)
}
