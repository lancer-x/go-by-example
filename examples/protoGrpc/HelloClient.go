package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())

	if err != nil {
		log.Fatal("err", err)
	}

	defer conn.Close()

	client := NewHelloServiceClient(conn)
	
	reply,err := client.Hello(context.Background(), &HelloMessage{
		Msg: "this is a test msg",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetMsg())
}
