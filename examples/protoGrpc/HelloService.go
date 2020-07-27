package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)
import "context"

type HelloServiceImpl struct {

}

func (p *HelloServiceImpl) Hello (ctx context.Context, args *HelloMessage) (*HelloMessage, error){
	reply := &HelloMessage{Msg:"hello:msg-->" + args.GetMsg()}
	return reply,nil
}

func main()  {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
