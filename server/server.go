package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"playground/golang-protobuf-example/protox"
)

type handler struct{}

func (s *handler) HandleOne(ctx context.Context, request *protox.ExampleRequest) (*protox.ExampleResponse, error) {
	response := &protox.ExampleResponse{
		Greeting: "Hi, I got your message: " + request.Message,
		Response: "We will definitely take a look at that",
	}

	return response, nil
}

func (s *handler) HandleTwo(ctx context.Context, request *protox.ExampleRequest) (*protox.ExampleResponse, error) {
	response := &protox.ExampleResponse{
		Greeting: "Hi, I got your message: " + request.Message,
		Response: "We are busy at the moment, please try again later",
	}

	return response, nil
}

func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	protox.RegisterRequestHandlerServer(srv, &handler{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
